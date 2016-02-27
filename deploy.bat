PATH=%PATH%;D:\cygwin64\bin
bash deploy-prepare.sh
D:\App\putty\pscp.exe -l pi -pw %RPIPASSWORD% content.tar 192.168.1.7:/home/pi/Dev/h_temp_mon/
del content.tar

