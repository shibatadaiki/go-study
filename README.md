# Golang Playground

## build
`docker-compose build`

## run
`docker-compose up`

## stop
`docker-compose down`

## debug

#### （`import "runtime"`をした後、コード上の任意の箇所で`runtime.Breakpoint()`を記述）

`docker exec -it go-study_app_1 bash`

`（仮想環境内で）dlv debug`
