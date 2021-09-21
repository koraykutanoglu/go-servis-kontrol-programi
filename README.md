## Go İle Linux Servislerinin Kontrolü Ve Otomatik Start Edilmesi

Bu github reposunda Linux üzerinde systemd olarak bu programı ayarlamayı ve ardından bu programı çalıştırarak servislerin 10 saniyede bir kontrolünü sağlayarak eğer bir servisin durması durumunda otomatik tekrardan başlatılmasını sağlayacaksınız. 

## Hangi Servisler Kontrol ediliyor?

Bu program üzerinde sırasıyla aşağıdaki servisler kontrol edilmektedir.

- docker.service
- sshd
- cron.service

## Linux Üzerine Programın Kurulumu

Sırasıyla SSH üzerinden aşağıdaki komutları girin

```
apt update
```

```
apt install golang -y
```

Programın Çalışacağı Dizini Ayarlayalım

```
mkdir -p /home/secops
```

```
mkdir -p /home/secops/app
```

```
cd /home/secops/app
```

```
git clone https://github.com/koraykutanoglu/go-servis-kontrol-programi
```

```
cd go-servis-kontrol-programi
```

Aşağıdaki dosyayı oluşturun

```
nano /etc/systemd/system/status-control.service
```

Aşağıdaki içeriği ekleyin

```
[Unit]
Description=Service Control
ConditionPathExists=/home/secops/app/go-servis-kontrol-programi
After=network.target


[Service]
Type=simple
User=root
Group=root
WorkingDirectory=/home/secops/app/go-servis-kontrol-programi
ExecStart=/usr/bin/go run .
Restart=on-failure
RestartSec=10
StandardOutput=syslog
StandardError=syslog
SyslogIdentifier=servicestatus


[Install]
WantedBy=multi-user.target
```

```
systemctl daemon-reload
```

```
sudo systemctl enable status-control.service
```

```
systemctl restart status-control.service
```

```
systemctl status status-control.service
```

Programın çıktısını aşağıdaki komutla kontrol ediyoruz

```
journalctl -f -u status-control.service
```

## Rsyslog Konfigürasyonu

```
mkdir /var/log/servicestatus
```

```
chown syslog:syslog /var/log/servicestatus
```

Aşağıdaki dosyayı düzenleyin

```
nano /etc/rsyslog.d/servicestatus.conf
```

Aşağıdaki içeriği ekleyin

```
if $programname == 'servicestatus' then /var/log/servicestatus/output.log
& stop
```

```
systemctl restart rsyslog.service
```

```
service servicestatus stop
```

```
service servicestatus restart
```

Programın çıktısını ve logları aşağıdaki komutlarla inceleyebilirsiniz

```
journalctl -b | grep servicestatus
```

```
watch tail -n 15 /var/log/servicestatus/output.log
```
