# Go Starter Kit

Repozytorium zostało stworzone w celu szybkiego rozpoczęcia pracy nad nowymi projektami w **GoLang** w systemie **Windows 10 64-bit**.

Osobiście korzystałem  z pomocy na stronach:
 - [Installing Go, Gocode, GDB and LiteIDE](https://www.ardanlabs.com/blog/2013/06/installing-go-gocode-gdb-and-liteide.html)
 - [Easy Go Programming Setup for Windows](http://www.wadewegner.com/2014/12/easy-go-programming-setup-for-windows/)

## Instalacja (ręczna):

**[1]** Przejdź na stronę [Downloads](https://golang.org/dl/) i pobierz plik [go1.13.5.windows-amd64.zip](https://dl.google.com/go/go1.13.5.windows-amd64.zip) dla **Windows 10 64-bit** lub inny odpowiedni dla swojego systemu operacyjnego.

**[2]** Stwórz na dysku twardym folder `C:\App\GoLang\` i wypakuj do niego zawartość pobranego pliku **ZIP**.

Masz uzyskać strukturę podobną do poniższej:

```
C:\App\GoLang\api\
C:\App\GoLang\bin\
...
C:\App\GoLang\pkg\
...
C:\App\GoLang\SECURITY.md
C:\App\GoLang\bin\VERSION
```

**[3]** Stwórz katalogi:

```
C:\App\GoLang\packages\
C:\App\GoLang\workspace\
```

**[4]** Następnie trzeba ustawić zmienne środowiskowe:

```
GOPATH      C:\App\GoLang\packages;C:\App\GoLang;C:\App\GoLang\workspace
GOROOT      C:\App\GoLang
GOTOOLDIR   C:\App\GoLang\pkg\tool\windows_amd64
```

**[5]** A do zmiennej **Path** trzeba dodać:

```
C:\App\GoLang\bin;C:\App\GoLang\packages\bin;C:\App\GoLang\workspace\bin
```

## Wtyczka dla VSC:

Wtyczka używana prze zemnie _( i działając :D )_ to:

```
Name: Go
Id: ms-vscode.go
Description: Rich Go language support for Visual Studio Code
Version: 0.11.9
Publisher: Microsoft
VS Marketplace Link: https://marketplace.visualstudio.com/items?itemName=ms-vscode.Go
```

## Komendy (polecenia):

Wszystkie komendy wykonujemy w głównym katalogu projektu!

 - Uruchamianie aplikacji:
	 - `go run ".\source"`
 - Budowanie pliku wykonywalnego _(EXE dla Windows)_:
	 - `go build -o ".\build\app.exe" ".\source"`