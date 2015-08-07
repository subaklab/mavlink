[![Build Status](https://travis-ci.org/mavlink/mavlink.svg?branch=master)](https://travis-ci.org/mavlink/mavlink)

## MAVLink ##

*   공식 사이트: http://mavlink.org
*   소스: [Mavlink Generator](https://github.com/mavlink/mavlink)
*   바이너리 (항상 master에서 최신 버전 유지):
  * [C/C++ header-only library](https://github.com/mavlink/c_library)
*   메일링 리스트: [Google Groups](http://groups.google.com/group/mavlink)

MAVLink -- Micro Air Vehicle Message Marshalling Library.

마이크로 비행 장치(Micro Air Vehicles) 사이나 그라운드 컨트롤 스테이션 간 가벼운 통신을 위한 라이브러리다. XML 파일내에 메시지를 정의할 수 있으며 따라서 적절한 이종의 언어에서  적절한 소스코드로 변환이 가능하다. 이런 XML 파일들을 다일렉트(dialect)라 부른다. 대부분은 `common.xml`에서 제공하는 *공통* 다일렉트를 기반으로 한다.

MAVLink 프로토콜은 byte-level 시리얼라이제이션을 수행하기에 radio modem 타입을 사용하는데 적합하다.

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

mavgen.py is a command-line interface for generating a language-specific MAVLink library. This is actually the backend used by `mavgenerate.py`. After the `mavlink` directory has been added to the Python path, it can be run by executing from the command line:

    $ python -m pymavlink.tools.mavgen

### Usage ###

Using the generated MAVLink dialect libraries varies depending on the language, with language-specific details below:

#### C ####
To use MAVLink, include the *mavlink.h* header file in your project:

    #include <mavlink.h>

Do not include the individual message files. In some cases you will have to add the main folder to the include search path as well. To be safe, we recommend these flags:

    $ gcc -I mavlink/include -I mavlink/include/<your message set, e.g. common>

The C MAVLink library utilizes a channels metaphor to allow for simultaneous processing of multiple MAVLink streams in the same program. It is therefore important to use the correct channel for each operation as all receiving and transmitting functions provided by MAVLink require a channel. If only one MAVLink stream exists, channel 0 should be used by using the `MAVLINK_COMM_0` constant.

##### Receiving ######
MAVLink reception is then done using `mavlink_helpers.h:mavlink_parse_char()`.

##### Transmitting #####
Transmitting can be done by using the `mavlink_msg_*_pack()` function, where one is defined for every message. The packed message can then be serialized with `mavlink_helpers.h:mavlink_msg_to_send_buffer()` and then writing the resultant byte array out over the appropriate serial interface.

It is possible to simplify the above by writing wrappers around the transmitting/receiving code. A multi-byte writing macro can be defined, `MAVLINK_SEND_UART_BYTES()`, or a single-byte function can be defined, `comm_send_ch()`, that wrap the low-level driver for transmitting the data. If this is done, `MAVLINK_USE_CONVENIENCE_FUNCTIONS` must be defined.

### Scripts/Examples ###
This MAVLink library also comes with supporting libraries and scripts for using, manipulating, and parsing MAVLink streams within the pymavlink, pymav
link/tools, and pymavlink/examples directories.

The scripts have the following requirements:
  * Python 2.7+
  * mavlink repository folder in `PYTHONPATH`
  * Write access to the entire `mavlink` folder.
  * Your dialect's XML file is in `message_definitions/*/`

Running these scripts can be done by running Python with the '-m' switch, which indicates that the given script exists on the PYTHONPATH. This is the proper way to run Python scripts that are part of a library as per PEP-328 (and the rejected PEP-3122). The following code runs `mavlogdump.py` in `/pymavlink/tools/` on the recorded MAVLink stream `test_run.mavlink` (other scripts in `/tools` and `/scripts` can be run in a similar fashion):

    $ python -m pymavlink.tools.mavlogdump test_run.mavlink

### License ###

MAVLink is licensed under the terms of the Lesser General Public License (version 3) of the Free Software Foundation (LGPLv3). The C-language version of MAVLink is a header-only library, and as such compiling an application with it is considered "using the library", not a derived work. MAVLink can therefore be used without limits in any closed-source application without publishing the source code of the closed-source application.

See the *COPYING* file for more info.

### Credits ###

&copy; 2009-2014 [Lorenz Meier](mailto:mail@qgroundcontrol.org)
