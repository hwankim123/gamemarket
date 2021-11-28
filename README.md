# Game Market
2021-2학기 FSSN 텀 프로젝트 : grpc를 사용한 게임 거래소 시스템 시뮬레이션

## 햇갈렸던 점
 - go mod init 시 생성되는 go.mod의 module path에 따라 package의 import 경로도 달라진다.
 
 ex) go/src/github.com/hwankim123에 example 하위 디렉토리를 생성하여 진행할 경우

 go mod init의 결과로 module github.com/hwankim123/example 과 같이 경로가 잡혔다면 import하려는 hwankim123 디렉토리 내의 package의 시작 경로도 github.com/이다.

 혹은 go mod init의 결과로 module example 과 같이 경로가 잡혔다면 import하려는 hwankim123 디렉토리 내의  package의 시작 경로도 example/이다.

 - 그리고 내가 만든 package를 import하려고 할 때 경로만 잘 설정하였다면 VSCode - GO extension의 auto import 기능으로 package를 자동으로 import해준다.(확실하지 않음. 구글링 필요) 그리고 경험상 import하려는 package에 naming을 따로 해주는 경우에는 수동으로 package 경로를 설정해줘야 하는 것 같다.
