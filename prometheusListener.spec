%ifarch aarch64
%global _arch aarch64
%global BuildArchitectures aarch64
%endif

%ifarch x86_64
%global _arch x86_64
%global BuildArchitectures x86_64
%endif

%define debug_package   %{nil}
%define _build_id_links none
%define _name prometheusListener
%define _prefix /opt
%define _version 1.04.00
%define _rel 0
%define _binaryname prometheusSDlistener

Name:       prometheusListener
Version:    %{_version}
Release:    %{_rel}
Summary:    Prometheus File-based Service Discovery listener

Group:      monitoring api
License:    GPL2.0
URL:        https://git.famillegratton.net:3000/monitoring/prometheusListener

Source0:    %{name}-%{_version}.tar.gz
#BuildArchitectures: x86_64
BuildRequires: gcc
#Requires:
#Obsoletes:

%description
Prometheus File-based Service Discovery listener

%prep
%autosetup

%build
cd %{_sourcedir}/%{_name}-%{_version}/src
PATH=$PATH:/opt/go/bin go build -o %{_sourcedir}/%{_binaryname} .
strip %{_sourcedir}/%{_binaryname}

%clean
rm -rf $RPM_BUILD_ROOT

%pre
getent group prometheus > /dev/null 2>&1 || groupadd prometheus
getent passwd prometheus > /dev/null 2>&1 || useradd -g prometheus -d /opt/prometheus -m -s /bin/bash prometheus > /dev/null
exit 0

%install
install -d %{buildroot}/opt/sbin
install -d %{buildroot}/etc/systemd/system/
install -Dpm 0755 %{_sourcedir}/%{_binaryname} %{buildroot}/opt/sbin/%{_binaryname}

%post
touch /etc/prometheusSDlistener.json
chown -R prometheus:prometheus /etc/prometheusSDlistener

cat << EOF > /etc/systemd/system/prometheusSDlistener.service
[Unit]
Description=Prometheus SD Listener Service
After=network.target

[Service]
User=prometheus
Group=prometheus
Type=simple
ExecStart=/opt/sbin/prometheusSDlistener
Restart=on-failure
RestartSec=10
# The following is for future use as the daemon does not log right now
#StandardOutput=append:/var/log/prometheusSDlistener.log
#StandardError=append:/var/log/prometheusSDlistener.err

[Install]
WantedBy=multi-user.target
EOF

chmod 644 /etc/systemd/system/prometheusSDlistener.service
systemctl daemon-reload

%preun
systemctl stop prometheusSDlistener
systemctl disable prometheusSDlistener

%postun
rm -f /etc/systemd/system/prometheusSDlistener.service
systemctl daemon-reload

%files
%defattr(-,root,root,-)
/opt/sbin/%{_binaryname}

%changelog
* Sun Sep 22 2024 RPM Builder <builder@famillegratton.net> 1.04.00-0
- Moved config file in another directory (jean-francois@famillegratton.net)

* Fri Sep 20 2024 RPM Builder <builder@famillegratton.net> 1.03.03-0
- Final script fix (alpine) (jean-francois@famillegratton.net)
- added runtime dependencies (jean-francois@famillegratton.net)
- More alpine build script fixes (jean-francois@famillegratton.net)
- Fixed alpine build scripts (jean-francois@famillegratton.net)
- Changed message in targets dir, version bump (jean-
  francois@famillegratton.net)
- Fixing DEB, now (jean-francois@famillegratton.net)
- Version bump for debian (builder@famillegratton.net)

* Wed Sep 18 2024 RPM Builder <builder@famillegratton.net> 1.03.02-0
- More path fixes (jean-francois@famillegratton.net)

* Wed Sep 18 2024 RPM Builder <builder@famillegratton.net> 1.03.01-0
- new package built with tito

* Wed Sep 18 2024 RPM Builder <builder@famillegratton.net> 1.03.00-0
- Version bump, forgotten in previous commit (jean-francois@famillegratton.net)
- Config file path fix (jean-francois@famillegratton.net)

