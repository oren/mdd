# Markdown driven development


Install silk

```bash
git clone https://github.com/matryer/silk
cd silk
go get
go install
```

Run this project

```
git clone https://github.com/oren/mdd
cd mdd
go build
./mdd
silk -silk.url="http://localhost:3000" api.md
```
