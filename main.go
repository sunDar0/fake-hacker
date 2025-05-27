package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"math/rand"
	"net"
	"net/http"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

//go:embed fake-hacker.html
var htmlContent string

//go:embed scenarios/*.json
var scenarioFiles embed.FS

type ScenarioLine struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type Scenario []ScenarioLine

var scenarios []Scenario

func init() {
	rand.Seed(time.Now().UnixNano())
	loadScenarios()
}

func loadScenarios() {
	// 내장된 시나리오 파일들 읽기
	entries, err := fs.ReadDir(scenarioFiles, "scenarios")
	if err != nil {
		log.Printf("시나리오 디렉토리 읽기 오류: %v", err)
		return
	}

	for _, entry := range entries {
		if !strings.HasSuffix(entry.Name(), ".json") {
			continue
		}

		filePath := "scenarios/" + entry.Name()
		data, err := scenarioFiles.ReadFile(filePath)
		if err != nil {
			log.Printf("파일 읽기 오류 %s: %v", filePath, err)
			continue
		}

		var scenario Scenario
		if err := json.Unmarshal(data, &scenario); err != nil {
			log.Printf("JSON 파싱 오류 %s: %v", filePath, err)
			continue
		}

		scenarios = append(scenarios, scenario)
	}

	log.Printf("%d개의 시나리오를 로드했습니다.", len(scenarios))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// 내장된 HTML 내용을 템플릿으로 파싱
	t, err := template.New("index").Parse(htmlContent)
	if err != nil {
		http.Error(w, "템플릿 파싱 오류: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t.Execute(w, nil)
}

func scenarioHandler(w http.ResponseWriter, r *http.Request) {
	if len(scenarios) == 0 {
		http.Error(w, "시나리오가 없습니다", http.StatusNotFound)
		return
	}

	// 랜덤 시나리오 선택
	randomIndex := rand.Intn(len(scenarios))
	selectedScenario := scenarios[randomIndex]

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(selectedScenario)
}

func scenarioListHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	response := map[string]interface{}{
		"count":     len(scenarios),
		"scenarios": scenarios,
	}
	
	json.NewEncoder(w).Encode(response)
}

func specificScenarioHandler(w http.ResponseWriter, r *http.Request) {
	// URL에서 시나리오 번호 추출
	path := strings.TrimPrefix(r.URL.Path, "/api/scenario/")
	index, err := strconv.Atoi(path)
	if err != nil {
		http.Error(w, "잘못된 시나리오 번호", http.StatusBadRequest)
		return
	}

	if index < 1 || index > len(scenarios) {
		http.Error(w, "시나리오 번호가 범위를 벗어났습니다", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(scenarios[index-1])
}

// 사용 가능한 포트를 찾는 함수
func findAvailablePort(startPort int) int {
	for port := startPort; port <= startPort+100; port++ {
		addr := fmt.Sprintf(":%d", port)
		listener, err := net.Listen("tcp", addr)
		if err == nil {
			listener.Close()
			return port
		}
	}
	return startPort // 기본값 반환
}

// 터미널 창을 최소화하는 함수
func minimizeWindow() {
	switch runtime.GOOS {
	case "windows":
		// Windows에서 현재 콘솔 창 최소화
		exec.Command("powershell", "-Command", "Add-Type -TypeDefinition 'using System; using System.Runtime.InteropServices; public class Win32 { [DllImport(\"user32.dll\")] public static extern bool ShowWindow(IntPtr hWnd, int nCmdShow); [DllImport(\"kernel32.dll\")] public static extern IntPtr GetConsoleWindow(); }'; $consolePtr = [Win32]::GetConsoleWindow(); [Win32]::ShowWindow($consolePtr, 2)").Start()
	case "darwin":
		// macOS에서 터미널 앱을 최소화 (여러 터미널 앱 지원)
		// 먼저 Terminal 앱 시도
		err := exec.Command("osascript", "-e", "tell application \"Terminal\" to set miniaturized of front window to true").Run()
		if err != nil {
			// Terminal이 실행 중이 아니면 iTerm2 시도
			err = exec.Command("osascript", "-e", "tell application \"iTerm2\" to tell current window to set miniaturized to true").Run()
			if err != nil {
				// 둘 다 안 되면 현재 활성 창 최소화
				exec.Command("osascript", "-e", "tell application \"System Events\" to tell process (name of first process whose frontmost is true) to set value of attribute \"AXMinimized\" of front window to true").Start()
			}
		}
	case "linux":
		// Linux에서는 wmctrl 사용 (설치되어 있다면)
		exec.Command("wmctrl", "-r", ":ACTIVE:", "-b", "add,hidden").Start()
	}
}

// 브라우저를 자동으로 여는 함수
func openBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("지원하지 않는 플랫폼입니다")
	}

	if err != nil {
		log.Printf("브라우저 열기 실패: %v", err)
		fmt.Printf("수동으로 브라우저에서 %s 를 열어주세요\n", url)
	}
}

func main() {
	// 라우트 설정
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/api/scenario", scenarioHandler)
	http.HandleFunc("/api/scenarios", scenarioListHandler)
	http.HandleFunc("/api/scenario/", specificScenarioHandler)

	// 사용 가능한 포트 찾기
	availablePort := findAvailablePort(8080)
	port := fmt.Sprintf(":%d", availablePort)
	url := fmt.Sprintf("http://localhost:%d", availablePort)
	
	fmt.Printf("🚀 Fake Hacker Terminal을 시작합니다...\n")
	fmt.Printf("🎬 %d개의 시나리오가 준비되었습니다\n", len(scenarios))
	fmt.Printf("🌐 포트 %d에서 서버를 시작합니다\n", availablePort)
	fmt.Printf("🔗 브라우저를 자동으로 열고 있습니다...\n")
	fmt.Printf("⚡ 종료하려면 Ctrl+C를 누르세요\n\n")

	// 서버를 고루틴으로 시작
	go func() {
		if err := http.ListenAndServe(port, nil); err != nil {
			log.Printf("서버 시작 실패: %v", err)
		}
	}()

	// 잠시 대기 후 브라우저 열기
	time.Sleep(500 * time.Millisecond)
	openBrowser(url)
	
	// 브라우저가 열린 후 창 최소화
	time.Sleep(1500 * time.Millisecond)
	minimizeWindow()

	// 메인 고루틴이 종료되지 않도록 대기
	select {}
} 