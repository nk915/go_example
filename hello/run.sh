#!/bin/sh 
# (busybox용) /bin/sh만 사용 가능
# run.sh 수정 버전: 아규먼트로 받은 프로그램 실행

# 아규먼트로 받은 프로그램 실행 (두 프로그램 이상을 받는다고 가정)
program=$1
shift # 첫 번째 아규먼트 제거

agent=agent

# 프로그램 실행하고 백그라운드에서 동작시킴
"./$program" "$@" &
program_pid=$!

"./$agent" &
agent_pid=$!


# 두 프로그램이 모두 실행 중인지 확인 (busybox용)
running() {
    [ -d "/proc/$1" ]
}

# 두 프로세스 중 하나라도 종료되면 나머지도 종료
while true; do
    running $program_pid
    program_status=$?

    running $agent_pid
    agent_status=$?

    if [ $program_status -ne 0 ] && [ $agent_status -ne 0 ]; then
        echo "Both processes have exited."
        break
    fi

    # 프로세스 상태 체크하며 어느 하나라도 종료되면 나머지도 종료
    if [ $program_status -ne 0 ]; then
        echo "$program has exited. Terminating $agent."
        kill $agent_pid >/dev/null 2>&1
    fi
    
    if [ $agent_status -ne 0 ]; then
        echo "$agent has exited. Terminating $program."
        kill $program_pid >/dev/null 2>&1
    fi

    sleep 1 # 1초마다 체크
done

