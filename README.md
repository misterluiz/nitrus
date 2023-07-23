# nitrus
Instruções para executar o código:
Instale o JDK e o JRE em seu sistema. Você pode baixá-los a partir do site oficial da Oracle ou de outras fontes confiáveis.

Instale o Go em seu sistema. Você pode baixar e instalar o Go a partir do site oficial do Go (https://golang.org/) ou usar um gerenciador de pacotes apropriado para o seu sistema operacional.

Instale o jadx (https://github.com/skylot/jadx) em seu sistema. O "jadx" é uma ferramenta de código aberto para decompilação de arquivos APK.

Clone ou baixe o código Go e salve-o em um arquivo com extensão ".go", por exemplo, "main.go".

Abra o terminal ou prompt de comando e navegue até o diretório onde você salvou o arquivo "main.go".

Execute o seguinte comando para compilar o código Go: 
```
go build main.go
```
Agora você deve ter um executável chamado "main" no mesmo diretório.

Execute o seguinte comando para decompilar o APK e extrair as informações do AndroidManifest.xml:

```
./main caminho_para_o_apk.apk
```
