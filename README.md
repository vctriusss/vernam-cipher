# Шифр Вернама

Программа для шифрования и расшифрования файлов с помощью Шифра Вернама (по модулю m)

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

### Режимы

Так как шифр Вернама зависит от используемого алфавита, то при шифровании и расшифровании его нужно указать. Делается это с помощью флагов в `[MODES]` в примерах ниже.

Используются следующие флаги:
- язык - `--eng` или `--rus`
- буквы в нижнем регистре (`-l` или`--lower`)
- буквы в верхнем регистре (`-u` или`--upper`)
- пробел (`-s` или `--space`)
- цифры (`d` или `--digits`)
- знаки пунктуации (`p` или `--punctuation`)

Флаги с однобуквенными алиасами можно объединять

Наример, чтобы зашифровать текст, в котором есть русские буквы обоих регистров, а также пробел и восклицательный знак, флаги указываются вот так: `--rus -lusp` или `--rus --lower --upper --space --punctuation`

### Шифрование
```bash
vernam encrypt -i <input-file> -o <output-file> -k <key-file> [MODES]
```

Также можно не задавать ключ, тогда он будет сгенерирован и записан в файл `key.txt`. Тогда запуск программы будет выглядеть так
```bash
vernam encrypt -i <input-file> -o <output-file> [MODES]
```

### Расшифрование
```bash
vernam decrypt -i <input-file> -o <output-file> -k <key-file> [MODES]
```

### Помощь
```bash
vernam help
vernam -h
```
