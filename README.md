# gZURL (Go Zabbix URL)
open winbox, ssh, vnc, rdp, telnet on the zabbix map

Forked from:
https://akmalov.com/blog/zabbix-open-winbox/

**Manual:**
1. Copy gZurl.exe -> C:\\@\

2. Place winbox.exe, vncviewer.exe, putty.exe -> C:\\@\

3. Run  gzurl.reg 

4. On Zabbix server add "gzurl" to valid_uri

edit defines.inc.php 
nano /usr/share/zabbix/include/defines.inc.php
Ctrl+W - find: ZBX_URI_VALID_SCHEMES - add gzurl - Save and Exit

**or add gzurl to:**

Administration -> General -> Other -> Valid URI schemes


5. Zabbix -> map -> host - add URLS format
gzurl://**xxx** ipaddress login password

Where **xxx**
* wbx - winbox.exe(ipAddr, login, password);

* rdp - mstsc(ipAddr);

* vnc - vncviewer.exe(ipAddr);

* ssh - putty.exe(ipAddr, login, password(optional));

* tln - putty.exe(ipAddr);

**examples:** 

gzurl://ssh myServerName root

gzurl://wbx 192.168.88.1 admin 12345

gzurl://tln 10.1.3.4

You can also use zabbix macro:

gzurl://wbx {HOST.IP} admin 12345

