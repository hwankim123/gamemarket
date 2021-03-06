# Game Market
## 1. 프로젝트 계획 의도
> 2021-2학기 FSSN(풀스택 네트워킹 서비스) 수업의 텀 프로젝트로, grpc를 사용한 게임 거래소 시스템을 시뮬레이션하는 프로젝트입니다. 이전 개인 프로젝트에서 HTTP 기반의 RESTful 서버를 개발해보았고, 새로운 네트워킹 방식을 경험해보고 싶어 gRPC 기술을 선택했습니다. 언어 또한 가볍게 경험해본다는 마음으로 Go언어를 선택했습니다.

------------
## 2. 요구사항
> ### 2.1. client
> > * client는 시간 상의 문제로 gRPC connection & stub을 생성 및 활용하는 코드에서 콘솔에 response받은 데이터를 바로 출력하는 형식으로 구현
> ### 2.2. server
> > * server는 gRPC의 네트워킹 패턴 중 가장 간단한 simple rpc 방식을 사용
> > * DB를 따로 적용하지 않았기 때문에 메모리에 임시로 데이터를 생성하는 방식을 사용
> ### 2.3. 기능
> > #### 2.3.1. 아이템 검색
> > > * 아이템 이름, 가격 범위, 옵션 범위(최대 3가지 옵션) 으로 검색 query를 구성
> > > * 아이템 검색의 결과로 아이템 id(식별자), 아이템 이름, 가격, 옵션을 출력
> > #### 2.3.2. 아이템 판매
> > > * 아이템 이름, 가격, 옵션으로 판매할 아이템의 정보를 입력
> > > * 아이템 판매의 결과로 사용자가 판매한 아이템의 id와 아이템의 정보를 출력
> > #### 2.3.3. 아이템 구매
> > > * 아이템 id를 입력해 구매할 아이템을 선택
> > > * 아이템 구매의 결과로 사용자가 구매한 아이템의 id와 아이템의 정보를 출력

------------
## 3. 햇갈렸던 점
> ### 3.1. import할 package의 경로 설정
> > * .proto 파일을 컴파일하여 생긴 go 코드와, 모듈화를 위해 코드를 분리하여 새로 생긴 go 코드를 main.go에 import하려고 할 때 경로가 잘 잡히지 않는 문제가 생겼다.
> > * go mod init 시 생성되는 go.mod의 module path에 따라 package의 import 경로도 달라진다.
> >
> > > * ex) go/src/github.com/hwankim123에 'example'이라는 하위 디렉토리를 생성할 경우
> > > 1) go mod init의 결과로 module github.com/hwankim123/example 과 같이 경로가 잡혔다면 import하려는 package의 경로의 시작 directory도 github.com/이다.
> > > 2) go mod init의 결과로 module example 과 같이 경로가 잡혔다면 import하려는 package의 경로의 시작 directory도 example/이다.
> >
> > * 내가 만든 package를 import하려고 할 때 경로만 잘 설정하였다면 VSCode - GO extension의 auto import 기능으로 package를 자동으로 import해준다.
> > * 그리고 경험상 import하려는 package에 naming을 따로 해주는 경우에는 수동으로 package 경로를 설정해줘야 하는 것 같다.
> >
> ### 3.2. context에 대한 이해도 X
> > * 이건 go언어를 접한지 얼마 되지 않아서 생긴 문제. 시간이 날 때 추가로 공부할 예정

------------
## 4. 느낀 점
> go 언어가 배우기 쉽다는 말에 현혹되어 급하게 gRPC - go 언어 프로젝트를 시작했습니다. 하지만 package의 경로를 잡지 못해 며칠간 삽질을 하고, 동작 원리를 제대로 이해하지 못한 채 마무리를 지어버리는 제 자신을 발견했습니다. 쉽든 어렵든, 뭐든 간에 천천히 깊게 공부해야 겠다고 느꼈습니다. go언어를 메인 언어로 삼을 예정은 아니지만, 사이드로 gRPC와 go를 공부할 예정입니다. 그래도 HTTP에서 벗어나 새로운 네트워킹 기술을 새로운 언어로 개발할 수 있어 좋은 경험이었습니다.
