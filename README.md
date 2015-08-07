[![Build Status](https://travis-ci.org/mavlink/mavlink.svg?branch=master)](https://travis-ci.org/mavlink/mavlink)

## MAVLink ##

*   공식 사이트: http://mavlink.org
*   소스: [Mavlink Generator](https://github.com/mavlink/mavlink)
*   바이너리 (항상 master에서 최신 버전 유지):
  * [C/C++ header-only library](https://github.com/mavlink/c_library)
*   메일링 리스트: [Google Groups](http://groups.google.com/group/mavlink)

MAVLink -- Micro Air Vehicle Message Marshalling Library.

마이크로 비행 장치(Micro Air Vehicles) 사이나 그라운드 컨트롤 스테이션 간 가벼운 통신을 위한 라이브러리다. XML 파일내에 메시지를 정의할 수 있으며 따라서 적절한 이종의 언어에서  적절한 소스코드로 변환이 가능하다. 이런 XML 파일들을 다일렉트(dialect)라 부른다. 대부분은 `common.xml`에서 제공하는 *공통* 다일렉트를 기반으로 한다.

MAVLink 프로토콜은 byte-level 직렬화를 수행하기에 radio modem 타입을 사용하는데 적합하다.

이 레파지토리는 거대한 파이썬 스크립트로 XML 파일을 언어 따른 라이브러리로 변환한다. MAVLink 데이터와 동작하는 예제와 유틸리티를 위해 추가적인 파이썬 스크립트도 있다. MAVLink 다일렉트를 위해 뿐만 아니라 이 스크립트는 파이썬 2.7 혹은 그 이상 버전을 필요로 한다.

2가지 MAVLink 프로토콜은 서로 호환되지 않는다는 점을 명심하자. : v0.9와 v1.0. [QGroundControl](https://github.com/mavlink/qgroundcontrol)을 포함한 대부분 프로그램은 이미 v1.0으로 넘어갔다. v0.9 프로토콜은 필요에 따라서 하위 호환성을 위해서만 사용된다.

### 요구사항 ###
  * Python 2.7+
    * Tkinter (만약 GUI 기능이 필요한 경우에)

### 설치 ###
  1. 사용자 쓰기 가능한 디렉토리로 Clone
  2. 레포지토리 디렉토리를 `PYTHONPATH`에 추가
  3. MAVLink 파서 파일을 생성. 다음 섹션에 있는 *AND/OR* 명령을 따라서 실행. Scripts/Examples 섹션에 설명된 helper 스크립트 포함.

### 특정 언어별 소스 파일 생성하기 ###

파이썬 스크립트를 커맨드 라인이나 GUI를 사용해서 특정 언어별 파일을 생성할 수 있다. 만약 다일렉트 XML 파일이 다른 XML 파일에 의존성을 가지는 경우 동일한 디렉토리에 위치해야 한다. 따라서 대부분 MAVLink 다일렉트는 **common** messageset에 의존하므로 여러분의 다일렉트를 `/message_definitions/v1.0/`에 있는 다른 것들과 함께 두는 것을 권장한다.

가능한 언어 :

  * C
  * C#
  * Java
  * JavaScript
  * Lua
  * Python, version 2.7+

#### GUI에서 (추천방법) ####

mavgenerate.py가 GUI 헤더 생성 도구이다. Tkinter가 필요하며 이는 윈도우 설치 파이썬에만 포함되어 있다. 따라서 윈도우가 아닌 플랫폼에서는 별도로 설치가 필요하다. 파이썬의 -m 인자를 사용하는 어디에서나 실행이 가능하다:

    $ python -m mavgenerate

#### 커맨드 라인에서 ####

mavgen.py는 커맨드 라인 인터페이스로 언어별 MAVLink 라이브러리를 생성한다. 이것은 실제로 백엔드로 `mavgenerate.py`에서 사용된다. `mavlink` 디렉토리가 파이썬 패스에 추가되고 나면 아래와 같은 커맨드 라인으로 실행할 수 있다:

    $ python -m pymavlink.tools.mavgen

### 사용 ###

생성된 MAVLink 다일렉트 라이브러리를 사용하면 언어에 따라서 다르다. 언어에 따른 상세내용은 아래와 같다:

#### C ####
MAVLink를 사용하기 위해 여러분의 프로젝트에서 *mavlink.h* 헤더 파일을 include한다.

    #include <mavlink.h>

개별 message 파일을 include하지마라. 어떤 경우에는 메인 폴더를 include 검색 패스에 추가해야만 한다. 안전하게 하려면 아래와 같은 flag를 추천한다 :

    $ gcc -I mavlink/include -I mavlink/include/<your message set, e.g. common>

C MAVLink 라이브러리는 채널 메터포를 사용해서 동일 프로그램에서 여러 MAVLink stream의 동시 처리를 허용한다. 따라서 각 연산에서 올바른 채널을 사용하는 것이 중요하다. 모든 주고 받는 기능은 채널을 획득한 MAVLink에서 제공한다. 만약 하나의 MAVLink stream만 존재한다면, 채널 0은 `MAVLINK_COMM_0` 상수를 이용해서 사용된다.

##### 수신(Receiving) ######
MAVLink 수신은 `mavlink_helpers.h:mavlink_parse_char()`을 이용해서 수행된다.

##### 송신(Transmitting) #####
전송은 `mavlink_msg_*_pack()` 함수를 이용해서 수행된다. 여기서 모든 메시지에 대해서 하나가 정의된다. 패킹된 메시지는 `mavlink_helpers.h:mavlink_msg_to_send_buffer()`를 통해서 직렬화가 된다. 그런 다음 결과 바이트 배열이 적절한 시리얼 인터페이스 위에 나오게 된다.

transmitting/receiving 코드를 wrapper를 작성해서 위에 것들을 단순화 시킬 수도 있다. 멀티-바이트 쓰기 매크로도 `MAVLINK_SEND_UART_BYTES()`와 같이 정의할 수 있다. 혹은 단일-바이트 함수는 `comm_send_ch()`와 같이 정의할 수 있다. 데이터를 전송하기 위해서 로우레벨 드라이버를 감쌀 수 있다. 만약 이렇게 되면 `MAVLINK_USE_CONVENIENCE_FUNCTIONS`는 반드시 정의해야만 한다.

### Scripts/Examples ###
MAVLink 라이브러리는 여러 라이브러리와 스크립트를 지원한다. pymavlink, pymav
link/tools, and pymavlink/examples 디렉토리들 내에서 MAVLink stream을 사용, 처리 그리고 파싱할 수 있다.

이 스크립트는 다음과 같은 요구사항을 가진다:
  * Python 2.7+
  * `PYTHONPATH`내에 mavlink 레파지토리 폴더
  * 전체 `mavlink` 폴더에 쓰기 접근
  * 여러분의 다일렉트 XML 파일은 `message_definitions/*/`에 존재


이런 스크립트를 실행하는 것은 '-m' 스위치로 파이썬을 실행함으로써 수행할 수 있다. 이것은 주어진 스크립트가 PYTHONPATH에 존재한다는 것을 뜻한다. PEP-328에 따른 라이브러리의 부분으로 파이썬 스크립트를 실행하기 위한 적절한 방식이다.(PEP-3122는 거부되었음) 녹화된 MAVLink stream `test_run.mavlink`에서 다음과 같은 코드는 `/pymavlink/tools/`에 있는 `mavlogdump.py`을 실행한다. (`/tools`과 `/scripts`에 있는 다른 스크립트는 유사한 방식으로 실행할 수 있다.):

    $ python -m pymavlink.tools.mavlogdump test_run.mavlink

### 라이센스 ###

MAVLink는 LGPLv3(the Lesser General Public License (version 3) of the Free Software Foundation)에 해당하는 라이센스 정책을 따른다. MAVLink의 C언어 버전은 헤더만 있는 라이브러리이며 이것을 가지고 어플리케이션을 컴파일하는 것은 "라이브러리를 사용했다"로 여긴다. MAVLink는 따라서 어떠한 공개하지 않는 소스 어플리케이션에서 제한없이 사용이 가능하다. 

보다 자세한 내용은 *COPYING* 파일을 참조하라.

### Credits ###

&copy; 2009-2014 [Lorenz Meier](mailto:mail@qgroundcontrol.org)
