# Morse · ... --- ...

[![CI](https://github.com/jzes/morse/actions/workflows/ci.yml/badge.svg)](https://github.com/jzes/morse/actions/workflows/ci.yml)
[![Release](https://img.shields.io/github/v/release/jzes/morse)](https://github.com/jzes/morse/releases/latest)
[![Go](https://img.shields.io/github/go-mod/go-version/jzes/morse)](https://go.dev/)

**De pontos e traços para palavras, direto no terminal.**

Um decoder de código Morse escrito em Go, com uma árvore binária no coração e zero dependências externas. Pequeno o suficiente para explorar o código inteiro; divertido o suficiente para começar com um `sos` e querer decifrar a próxima mensagem.

```text
$ morse '.... . .-.. .-.. ---  .-- --- .-. .-.. -..'
hello world
```

## 💡 De um vídeo para o terminal

O projeto nasceu da curiosidade ao assistir a [este short no YouTube](https://www.youtube.com/shorts/oMheP7qCmyQ), que mostra um dispositivo decodificando código Morse. A ideia virou um exercício: como representar pontos e traços em código e transformá-los em letras?

A resposta aqui é uma árvore: cada ponto segue para um lado, cada traço para o outro, e o caminho percorrido revela a letra.

## 🚀 Instalação

### Homebrew (macOS e Linux)

```sh
brew install jzes/tap/morse
morse '... --- ...'
```

O [tap oficial do projeto](https://github.com/jzes/homebrew-tap) instala os binários publicados nas releases, com verificação SHA-256. Há suporte a macOS e Linux, em `amd64` e `arm64`, sem precisar instalar Go.

### Baixar um binário

Na [página de releases](https://github.com/jzes/morse/releases), baixe o arquivo correspondente ao seu sistema e arquitetura e extraia seu conteúdo:

| Sistema | Arquiteturas | Arquivo |
| --- | --- | --- |
| Linux | `amd64`, `arm64` | `.tar.gz` |
| macOS (`darwin`) | `amd64` (Intel), `arm64` (Apple Silicon) | `.tar.gz` |
| Windows | `amd64`, `arm64` | `.zip` |

Execute `./morse` no Linux/macOS ou `.\morse.exe` no PowerShell. Para usar apenas `morse`, coloque o executável em um diretório do seu `PATH`.

Cada release inclui um `checksums.txt` com os hashes SHA-256 dos arquivos.

### Instalar com Go

Com **Go 1.26.4 ou superior**:

```sh
go install github.com/jzes/morse@latest
morse '... --- ...'
```

O executável vai para `$(go env GOPATH)/bin`, ou para `GOBIN` quando configurado. Esse diretório precisa estar no seu `PATH`.

### Rodar a partir do código

```sh
git clone https://github.com/jzes/morse.git
cd morse
go run . '... --- ...'
```

## 📡 Como usar

### Uma mensagem como argumento

Use aspas para que toda a mensagem seja passada como um único argumento:

```sh
morse '... --- ...'
# sos

morse '--. ---  .-. --- -.-. -.- ...'
# go rocks
```

### Entrada padrão: pipes e arquivos

Sem argumento, o programa lê de `stdin` até o fim da entrada:

```sh
printf '%s' '.... . .-.. .-.. ---  .-- --- .-. .-.. -..' | morse
# hello world

morse < mensagem.txt
```

### Formato da mensagem

| Entrada | Significado |
| --- | --- |
| `.` | Ponto |
| `-` | Traço |
| Um espaço | Separa letras |
| Dois espaços | Separam palavras |

O alfabeto implementado contém **as 26 letras de A a Z**, e a saída é em letras minúsculas, seguida de uma quebra de linha.

> Use o formato acima para obter uma decodificação correta. Números e pontuação não estão implementados. A barra `/` não é um separador de palavras: caracteres diferentes de ponto, traço e espaço são ignorados, inclusive quebras de linha. A validação de sequências Morse inválidas ainda é limitada.

## 🌳 Como funciona

Cada nó guarda uma letra e dois caminhos possíveis: `Dot` e `Dash`.

```text
                 início
                /      \
             . /        \ -
              E          T
             / \        / \
          . /   \ -  . /   \ -
           I     A    N     M
          / \   / \  / \   / \
         S   U R   W D   K G   O
         ⋮   ⋮ ⋮   ⋮ ⋮   ⋮ ⋮
```

Para decodificar `...`, por exemplo:

1. Começa na raiz.
2. O primeiro `.` leva a **E**.
3. O segundo `.` leva a **I**.
4. O terceiro `.` leva a **S**.
5. Um espaço encerra a letra e reinicia o percurso na raiz. O fim da entrada também encerra a última letra.

O decoder recebe um `io.Reader` e um `io.Writer`, permitindo usar a mesma lógica com entrada do terminal, arquivos ou buffers. Atualmente, a entrada é lida inteira em memória antes da decodificação.

### Estrutura do projeto

```text
.
├── main.go                # CLI: argumento ou stdin → stdout
├── alphabet/
│   └── alphabet.go        # Árvore com as 26 letras
├── decoder/
│   └── decoder.go         # Percorre a árvore e decodifica a entrada
└── .github/workflows/
    ├── ci.yml             # Formatação, análise, testes, build e exemplos
    └── release.yml        # Binários multiplataforma e GitHub Release
```

## 🛠️ Desenvolvimento

```sh
go fmt ./...
go vet ./...
go test ./...
go build -o morse .
./morse '... --- ...'
```

Os testes do pacote `decoder` cobrem o alfabeto completo, mensagens, separadores, entrada vazia, uma árvore personalizada e falhas de leitura. Além de executar `go test`, o CI verifica a formatação, roda `go vet`, compila o programa e confere exemplos reais da CLI por argumento e por `stdin`. Ele é executado em pushes para `main` e em pull requests.

## 📦 Publicando uma versão

O release é disparado pelo envio de uma **tag com prefixo `v`**. Depois de integrar as alterações à `main` e atualizar seu checkout local:

```sh
git switch main
git pull --ff-only
git tag -a v0.1.0 -m "Release v0.1.0"
git push origin v0.1.0
```

Troque `v0.1.0` pela próxima versão, seguindo `vMAJOR.MINOR.PATCH`. O workflow:

1. Executa as mesmas verificações do CI.
2. Compila para Linux, macOS e Windows, em `amd64` e `arm64`.
3. Empacota os executáveis com este README e a licença MIT e gera checksums SHA-256.
4. Publica uma GitHub Release com notas geradas automaticamente e os arquivos para download.

Tags como `v0.2.0-rc.1` geram uma **pré-release**. A versão é escolhida na tag; não há incremento automático. A publicação usa o `GITHUB_TOKEN` fornecido pelo Actions, sem precisar cadastrar um token pessoal.

Após publicar uma versão estável, atualize a versão, as URLs e os checksums da fórmula `Formula/morse.rb` no [repositório do tap](https://github.com/jzes/homebrew-tap). A atualização do tap é manual.

## 📄 Licença

Distribuído sob a [licença MIT](LICENSE).

---

**Uma árvore, dois sinais e muita curiosidade.** `... . -- .--. .-. .  .- .--. .-. . -. -.. . -. -.. ---`
