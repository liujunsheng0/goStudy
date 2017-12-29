package main

import (
	"fmt"
	"os"
)

/*
Linux操作系统每个进程都有一个父进程, 所有的进程共同则乘客一个树状结构.
内核启动进程作为root, 负责系统初始化操作, 如果一个进程先于它的子进程结束, 那么这些子进程会被root进程收养, 直接成为它的子进程
进程描述符(PID):记录每个进程的优先级, 状态, 访问权限, 虚拟地址等, (同一时刻, 进程号唯一, 利用这个可做很多)
Linux可凭借CPU快速地在多个进程间切换, 这就是进程的上下文切换, 无论切换速度如何, 同一时刻正在运行的进程仅有一个
*/

func getPid()  {
	fmt.Println("pid = ", os.Getpid()) // 进程号, 用来标识进程的数字, 内核可以快速的将进程id->进程描述符
	fmt.Println("父进程 id = ", os.Getppid())
}
func main() {
	getPid()
}