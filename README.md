# Obfuscate Command Tool

This incredible tool will turn transform a valid terminal command into an illegible jumble of base64 - which will still execute! 😮💥🎉

> What's the catch?

> How much will this marvel of modern technology set me back?

> This is surely the client API for a SaaS subscription right?

You've probably had all of the above thoughts. I don't blame you. I hope you're as surprised as me to realise that I'm offering this 100% free, forever.

## Installation
```
go get -u github.com/alexdcox/obfuscatecmd
cd $GOPATH/src/$_
go install
```

## Usage

```
obfuscatecmd help
```
> NAME:  
   > obfuscatecmd - Make an otherwise readable bash command completely illegible and many times the character length.  
>   
> USAGE:  
   > obfuscatecmd [command options] '[bash commands]'  
>   
> COMMANDS:  
     help, h  Shows a list of commands or help for one command  
>   
> OPTIONS:  
   > --iterations NUMBER, -i NUMBER  Determines the NUMBER of times to obfuscate and pack the provided command. (default: 5)  
   > --help, -h                      show help  
   > --version, -v                   print the version  
   
```
obfuscatecmd 'echo "amazing 🦆!"' > executeme.sh && chmod +x executeme.sh && ./$_
```

> amazing 🦆!

If you're not convinced by now...

```
cat executeme.sh
```
