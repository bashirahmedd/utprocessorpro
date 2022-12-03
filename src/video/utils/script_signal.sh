#!/bin/bash

# shutdown-now signal is higher in precedence, shutdown-now is not recommended
# shutdown is lower in precedence, shutdown is preferred
# exit is the lowest in precedence, 

fn_process_signal() 
{
    declare -a signal_file=(
        "./signal/shutdown-now" 
        "./signal/shutdown" 
        "./signal/exit"
        )
    msg_delay=5

    for signal in "${signal_file[@]}"
    do
        if [[ -f "$signal" && "$signal" == "${signal_file[0]}" ]]; then
            fn_say "System shutdown called."
            fn_say "System will proceed to shutdown immediately."
            sleep $msg_delay
            shutdown now
            exit 0
        elif [[ -f "$signal" && "$signal" == "${signal_file[1]}" ]]; then
            fn_say "System shutdown called."
            fn_say "System will proceed to shutdown in 120 seconds."
            sleep $msg_delay
            shutdown +2 Shutting down in 2 minutes!
            exit 0
        elif [[ -f "$signal" && "$signal" == "${signal_file[2]}" ]]; then
            fn_say "Stop download called."
            fn_say "Please do manual merge of input_id and try_again files."
            exit 0
        fi     
    done
}


# Test Calls for the above functions
#fn_exit_signal $*
#fn_exit_signal "Hello Hello, how are you?"
#./script_exit.sh "EB is naughty boy. EB is bad boy, he is touching computer. He is disturbing others. Computer is angry. Computer will slap EB on the cheeks. "