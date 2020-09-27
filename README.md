# Golang Playground

## build
`docker-compose build`

## run
`docker-compose up`

## stop
`docker-compose down`

## debug

（`import "runtime"`をした後、コード上の任意の箇所で`runtime.Breakpoint()`を記述）

`（docker-compose upされている状態で）docker exec -it go-study_app_1 bash`

`（仮想環境内で）dlv debug`

## doc

`docker run --rm -p 127.0.0.1:4000:4000 go-study_app godoc -http 0.0.0.0:4000`

## REPL

`（docker-compose upされている状態で）docker exec -it go-study_app_1 bash`

`（仮想環境内で）gore`
