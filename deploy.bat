bash deploy-prepare.sh
pscp.exe -l pi -pw %RPIPASSWORD% content.tar 192.168.1.7:/home/pi/Dev/h_temp_mon/
del content.tar

