FOR /F %%a IN ('git describe') DO (
set ver=%%a
)
@echo package main>BuildVersion.go
@echo const BuildVersion = "%ver%">>BuildVersion.go
