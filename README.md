### psgrep

grep ps command

```
$ ./psgrep nginx
UID        PID  PPID  C STIME TTY          TIME CMD
root     21389     1  0 май08 ?     00:00:00 nginx: master process /usr/sbin/nginx -g daemon on; master_process on;
www-data 21390 21389  0 май08 ?     00:00:01 nginx: worker process
$ PSGREP_OPTS="aux" ./psgrep nginx
USER       PID %CPU %MEM    VSZ   RSS TTY      STAT START   TIME COMMAND
root     21389  0.0  0.6 101560  3360 ?        Ss   май08   0:00 nginx: master process /usr/sbin/nginx -g daemon on; master_process on;
www-data 21390  0.0  1.4 102668  7148 ?        S    май08   0:01 nginx: worker process
```
