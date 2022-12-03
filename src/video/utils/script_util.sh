#!/bin/bash

fn_process_fsize() 
{
    echo "fn_process_fsize :""$1"

    #curr_file_sz=`youtube-dl -f 22/18/17 $1 -j | jq .filesize`
    #echo "Target raw file size: ""$curr_file_sz"
    #fl_sz=`numfmt --to=iec "$curr_file_sz"`
    #echo "Target file size: ""$fl_sz" 
    # if echo "$curr_file_sz" | grep -qE '^[0-9]+$'; then
    #     session_dl_sz="$(($session_dl_sz+$curr_file_sz))"
    # fi
}

fn_validate_file()
{

    local epoc_id=`date +%s`
    local validate_log="./log/""$epoc_id""_validate_id.log"
    local session_bkup_ids="$2"
    local target="$1"

    while read -r fid; do        
        echo $fid
        fname=`ls "$target"|grep -i "$fid"".mp4$" 2> /dev/null`
        if [[ $? -ne 0 ]];then
            echo $line >>  $validate_log
            echo "Please validate log, $validate_log"
        fi
    done <<< "`cat $session_bkup_ids`"
}

fn_remove_partial()
{
    echo "fn_remove_partial"
}

#fn_validate_file "/home/naji/Downloads/temp/ytdown/" "./log/1664218266_backup_video_id.log"     