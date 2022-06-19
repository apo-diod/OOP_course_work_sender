package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"
)

const SWD = "../OOP_course_work_modules/sender/"

func useModule(id string, data string) {
	log.Println("executing", id, "with", data)
	cmd := exec.Command("./venv/Scripts/python.exe", "script.py", data)
	cmd.Env = os.Environ()
	cmd.Dir = SWD + id
	cmd.Run()
}

func newModule(mtype string, settings string) string {
	if mtype == "save_to_file" {
		return newSaveToFile(settings)
	}
	if mtype == "send_request" {
		return newSendRequest(settings)
	}
	return "0"
}

func newSaveToFile(settings string) string {
	var set map[string]interface{}
	json.Unmarshal([]byte(settings), &set)
	path := set["path"].(string)
	id := RandStringRunes(16)
	os.Mkdir(SWD+id, 777)
	cmdvenv := exec.Command(SWD+"save_to_file/venv/Scripts/python.exe", "-m", "venv", "./venv")
	cmdvenv.Env = os.Environ()
	cmdvenv.Dir = SWD + id
	cmdvenv.Run()
	scriptbts, _ := os.ReadFile(SWD + "save_to_file/script.py")
	script := string(scriptbts)
	script = strings.Replace(script, "$", path, 1)
	f, _ := os.Create(SWD + id + "/script.py")
	f.WriteString(script)
	f.Close()
	cmdvenv.Wait()
	out, _ := cmdvenv.CombinedOutput()
	log.Println(string(out))
	return id
}

func newSendRequest(settings string) string {
	return "0"
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandStringRunes(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
