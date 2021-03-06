package main

import (
	"os"
	"os/exec"
	"strconv"
)

func main() {
	// Torに関する点
	// socksport でTor socksプロキシのポートを指定する。（使用しないがデフォルト設定だとポート重複で起動不可なため指定）
	// ORPortはTorのルーティングに使われるポート。指定するとリレーノードになる。
	// ControlPortは立ち上がったとにいろいろ追加でコマンドを送信できるポート。
	// Exitpolicy reject *:*とExitRelay 0入れてExit Nodeにならない。
	// DirPort はTorのディレクトリを通知するためのポート。
	// DataDirectoryはキャッシュデータなどいろいろなデータを格納するディレクトリ。
	// 	1つのプロセスにつき1つのディレクトリが必要。
	// NicknameでFITACJPという名前でリレーする。

	// 最大4999プレセス立てられる。
	maxProcess := 20
	i := maxProcess
	for maxProcess >= 0{
		maxProcess--
		// ポート番号生成
		socksPort := strconv.Itoa(0)
		controlPort := strconv.Itoa(40000 + maxProcess)
		orPort := strconv.Itoa(50000 + maxProcess)
		dirPort := strconv.Itoa(55000 + maxProcess)
		nickname := "FitAcJp" + strconv.Itoa(maxProcess)
		directoryName := "DataDirectory/" + strconv.Itoa(maxProcess)
		// データディレクトリ作成（ダブったらTorが起動できない）
		if err := os.MkdirAll(directoryName, 0777); err != nil {
			panic(err)
		}
		exec.Command("'ls'").Start()
		// tor -socksport 9050 -orport 9051 -exitrelay 0 -nickname FitAcJp -contactinfo s15a1302@bene.fit.ac.jp -exitpolicy "reject *:*" -dirport 9052
		err := exec.Command("tor",
			"-socksport", socksPort,
			"-orport",orPort,
			"-dirport",dirPort,
			"-controlport", controlPort, // 使わないのでコメントアウト
			"-exitrelay","0",
			"-exitpolicy",`reject *:*`,
			"-datadirectory", directoryName,
			"-MaxMemInQueues", strconv.Itoa(6144 / i) + "MB",
			"-contactinfo","s15a1302@bene.fit.ac.jp",
			"-nickname",nickname).Start()
				//">>","test.log").Run()
				//">> " + "log/" + strconv.Itoa(maxProcess) + ".log").Start()
		if err !=nil {
			panic(err)
		}
	}
}
