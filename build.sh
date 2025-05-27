#!/bin/bash

# 빌드 스크립트 - Fake Hacker Terminal
echo "🚀 Fake Hacker Terminal 빌드를 시작합니다..."

# output 디렉토리 생성
mkdir -p output

# 현재 플랫폼 정보 출력
echo "📋 빌드 대상 플랫폼:"
echo "   - macOS (darwin/amd64)"
echo "   - macOS Apple Silicon (darwin/arm64)"
echo "   - Windows (windows/amd64)"
echo "   - Linux (linux/amd64)"
echo ""

# macOS Intel
echo "🍎 macOS Intel 빌드 중..."
GOOS=darwin GOARCH=amd64 go build -o output/fake-hacker-macos-intel main.go
if [ $? -eq 0 ]; then
    echo "✅ macOS Intel 빌드 완료"
else
    echo "❌ macOS Intel 빌드 실패"
fi

# macOS Apple Silicon
echo "🍎 macOS Apple Silicon 빌드 중..."
GOOS=darwin GOARCH=arm64 go build -o output/fake-hacker-macos-arm64 main.go
if [ $? -eq 0 ]; then
    echo "✅ macOS Apple Silicon 빌드 완료"
else
    echo "❌ macOS Apple Silicon 빌드 실패"
fi

# Windows
echo "🪟 Windows 빌드 중..."
GOOS=windows GOARCH=amd64 go build -o output/fake-hacker-windows.exe main.go
if [ $? -eq 0 ]; then
    echo "✅ Windows 빌드 완료"
else
    echo "❌ Windows 빌드 실패"
fi

# Linux
echo "🐧 Linux 빌드 중..."
GOOS=linux GOARCH=amd64 go build -o output/fake-hacker-linux main.go
if [ $? -eq 0 ]; then
    echo "✅ Linux 빌드 완료"
else
    echo "❌ Linux 빌드 실패"
fi

echo ""
echo "📦 빌드 결과:"
ls -la output/

echo ""
echo "🎉 빌드가 완료되었습니다!"
echo ""
echo "📋 사용법:"
echo "   macOS Intel:        ./output/fake-hacker-macos-intel"
echo "   macOS Apple Silicon: ./output/fake-hacker-macos-arm64"
echo "   Windows:            output/fake-hacker-windows.exe"
echo "   Linux:              ./output/fake-hacker-linux"
echo ""
echo "✅ 모든 파일이 실행파일에 내장되어 있어 어디서든 독립적으로 실행 가능합니다!"
echo "🚀 실행파일을 더블클릭하면 자동으로 브라우저가 열리고 해커 화면이 시작됩니다." 