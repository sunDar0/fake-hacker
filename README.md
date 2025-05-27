# Fake Hacker Terminal - Go

해커 영화나 드라마에서 볼 수 있는 터미널 화면을 시뮬레이션하는 Go 웹 애플리케이션입니다. 실제 해킹 도구들의 명령어와 출력을 리얼하게 재현하여 타이핑 효과와 함께 보여줍니다.

## 🎯 주요 기능

### 1. 리얼한 터미널 시뮬레이션
- **검은 배경 + 녹색 텍스트**: 클래식한 해커 터미널 스타일
- **깜빡이는 커서**: 실제 터미널처럼 커서가 깜빡임
- **타이핑 효과**: 사용자 입력을 실제로 타이핑하는 것처럼 한 글자씩 표시
- **프롬프트 구분**: 입력 명령어는 `$` 프롬프트와 함께 표시

### 2. 동적 시나리오 시스템
- **랜덤 시나리오 선택**: 페이지를 새로고침할 때마다 다른 시나리오 실행
- **JSON 기반 시나리오**: 쉽게 편집 가능한 JSON 형식으로 시나리오 관리
- **확장 가능**: 새로운 시나리오 파일을 추가하면 자동으로 인식

### 3. 현실적인 타이밍
- **가변 타이핑 속도**: 15-80ms 사이의 랜덤한 타이핑 속도
- **지능적 딜레이**: 
  - 명령어 입력 완료 후 → 시스템 응답: 300ms (처리 시간)
  - 시스템 응답 완료 후 → 다음 명령어: 800ms (사용자 생각 시간)
  - 시스템 응답 간: 100ms

## 📁 프로젝트 구조

```
fake-hacker/
├── main.go               # Go 웹 서버 메인 파일
├── go.mod                # Go 모듈 파일
├── build.sh              # 크로스 플랫폼 빌드 스크립트
├── scenarios/            # 시나리오 JSON 파일들
│   ├── scenario-1.json   # 웹 애플리케이션 침투 테스트
│   ├── scenario-2.json   # 네트워크 스캔 및 분석
│   └── scenario-3.json   # 시스템 침투 및 권한 상승
├── output/               # 빌드된 실행 파일들
│   ├── fake-hacker-macos-intel
│   ├── fake-hacker-macos-arm64
│   ├── fake-hacker-windows.exe
│   └── fake-hacker-linux
├── fake-hacker.html      # HTML 템플릿 (main.go에 내장됨)
└── README.md
```

## 🎬 포함된 시나리오

### 시나리오 1: 웹 애플리케이션 침투 테스트
- 네트워크 스캔 (`nmap`, `ifconfig`)
- 웹 서버 정보 수집 (`curl`, `nikto`)
- 디렉토리 브루트포스 (`dirb`)
- SQL 인젝션 공격 (`sqlmap`)
- 데이터베이스 덤프 및 패스워드 크래킹

### 시나리오 2: 네트워크 분석
- 패킷 캡처 및 분석 (`wireshark`, `tcpdump`)
- 네트워크 서비스 스캔 (`netstat`, `ss`)
- 트래픽 모니터링

### 시나리오 3: 시스템 침투
- SSH 접속 및 권한 상승
- 시스템 파일 접근 (`/etc/shadow`)
- 사용자 권한 확인 (`sudo -l`)

## 🚀 사용 방법

### 방법 1: 미리 빌드된 실행 파일 사용 (권장)
```bash
# 크로스 플랫폼 빌드
./build.sh

# 플랫폼별 실행
# macOS Intel
./output/fake-hacker-macos-intel

# macOS Apple Silicon  
./output/fake-hacker-macos-arm64

# Windows
output/fake-hacker-windows.exe

# Linux
./output/fake-hacker-linux
```

### 방법 2: Go 소스코드로 직접 실행
```bash
# 서버 실행
go run main.go

# 또는 빌드 후 실행
go build -o fake-hacker
./fake-hacker
```

### 자동 실행 기능
- **자동 포트 감지**: 8080 포트가 사용 중이면 자동으로 다른 포트 사용
- **자동 브라우저 열기**: 실행하면 자동으로 기본 브라우저가 열림
- **터미널 최소화**: 브라우저가 열린 후 터미널 창이 자동으로 최소화됨

### API 엔드포인트
- `GET /` - 메인 터미널 페이지
- `GET /api/scenario` - 랜덤 시나리오 가져오기
- `GET /api/scenarios` - 모든 시나리오 목록
- `GET /api/scenario/{번호}` - 특정 시나리오 가져오기

### 실행 예시
```bash
# 실행 파일 다운로드 후
chmod +x fake-hacker-macos-arm64  # macOS/Linux의 경우
./fake-hacker-macos-arm64

# 출력 예시:
# 🚀 Fake Hacker Terminal을 시작합니다...
# 🎬 3개의 시나리오가 준비되었습니다
# 🌐 포트 8080에서 서버를 시작합니다
# 🔗 브라우저를 자동으로 열고 있습니다...
```

## 🛠 시나리오 추가/편집

### 새 시나리오 추가
1. `scenarios/` 폴더에 `scenario-4.json` 파일 생성
2. 다음 형식으로 작성:

```json
[
    {
        "type": "input",
        "text": "명령어"
    },
    {
        "type": "output", 
        "text": "출력 결과"
    }
]
```

### 시나리오 타입
- **`input`**: 사용자가 입력하는 명령어 (타이핑 효과 적용)
- **`output`**: 시스템 출력 (즉시 표시)

## 🎨 스타일 커스터마이징

CSS 변수를 수정하여 터미널 스타일을 변경할 수 있습니다:

```css
/* 색상 변경 */
background-color: #000000;  /* 배경색 */
color: #00FF00;             /* 텍스트 색상 */

/* 타이핑 속도 조정 */
const typingSpeedMin = 15;  /* 최소 속도 (ms) */
const typingSpeedMax = 80;  /* 최대 속도 (ms) */

/* 딜레이 시간 조정 */
const inputToOutputDelay = 300;   /* 명령어 → 응답 */
const outputToInputDelay = 800;   /* 응답 → 다음 명령어 */
const outputToOutputDelay = 100;  /* 응답 간 딜레이 */
```

## 🔧 기술 스택

- **Go 1.21+**: 백엔드 웹 서버
- **HTML5**: 기본 구조 (템플릿으로 내장)
- **CSS3**: 터미널 스타일링 및 애니메이션
- **Vanilla JavaScript**: 타이핑 효과 및 시나리오 로직
- **JSON**: 시나리오 데이터 저장
- **Go net/http**: 웹 서버 및 API
- **Go html/template**: HTML 템플릿 렌더링

## 📝 주의사항

- 이 프로젝트는 **교육 및 데모 목적**으로만 사용해야 합니다
- 실제 해킹 도구나 공격 기능은 포함되어 있지 않습니다
- 모든 명령어와 출력은 시뮬레이션된 가짜 데이터입니다

## ✨ Go 버전의 장점

- **단일 바이너리**: 의존성 없이 하나의 실행 파일로 배포
- **내장 웹 서버**: 별도 서버 설정 불필요
- **크로스 플랫폼**: Windows, macOS (Intel/ARM), Linux 지원
- **자동화 기능**: 포트 감지, 브라우저 열기, 터미널 최소화
- **빠른 시작**: 실행 파일 더블클릭만으로 즉시 시작
- **API 제공**: RESTful API로 시나리오 데이터 접근 가능
- **메모리 효율**: Go의 효율적인 메모리 관리
- **내장 리소스**: HTML, CSS, JS, 시나리오 파일 모두 실행파일에 포함

## 🔧 빌드 및 배포

### 크로스 플랫폼 빌드
```bash
# 모든 플랫폼용 빌드
./build.sh

# 개별 플랫폼 빌드
GOOS=darwin GOARCH=amd64 go build -o fake-hacker-macos-intel main.go
GOOS=darwin GOARCH=arm64 go build -o fake-hacker-macos-arm64 main.go
GOOS=windows GOARCH=amd64 go build -o fake-hacker-windows.exe main.go
GOOS=linux GOARCH=amd64 go build -o fake-hacker-linux main.go
```

### 배포 방법
1. **단일 파일 배포**: 실행 파일 하나만 복사하면 완료
2. **의존성 없음**: Go 런타임이나 추가 라이브러리 설치 불필요
3. **즉시 실행**: 더블클릭하면 자동으로 브라우저가 열리며 시작

## 🎯 활용 사례

- 해커 영화/드라마 촬영 소품
- 사이버 보안 교육 자료
- 프레젠테이션 데모
- 웹 개발 포트폴리오
- 재미있는 데스크톱 배경화면
- 서버 배포용 애플리케이션
- 크로스 플랫폼 데모 애플리케이션

---

**⚠️ 면책 조항**: 이 프로젝트는 순수히 교육 및 엔터테인먼트 목적으로 제작되었습니다. 실제 해킹이나 불법적인 활동에 사용해서는 안 됩니다. 