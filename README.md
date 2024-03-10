# Шифр Вернама

Программа для шифрования и расшифрования файлов с помощью Шифра Вернама

## Install

```bash
go install github.com/vctriusss/vernam-cipher/cmd/vernam@latest
```

or

```bash
git clone github.com/vctriusss/vernam-cipher.git
cd vernam-cipher
go build cmd/vernam/vernam.go
```

## Usage

### Шифрование
```bash
vernam encrypt -i <input-file> -o <output-file> -k <key-file>
```

Также можно не задавать ключ, тогда он будет сгенерирован и записан в файл `key.txt`. Тогда запуск программы будет выглядеть так
```bash
vernam encrypt -i <input-file> -o <output-file>
```

### Расшифрование
```bash
vernam decrypt -i <input-file> -o <output-file> -k <key-file>
```

### Помощь
```bash
vernam help
vernam -h
```
