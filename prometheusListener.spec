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
%define _version 1.02.00
%define _rel 0
%define _binaryname prometheusSDlistener

Name:       prometheusListener
Version:    %{_version}
Release:    %{_rel}
Summary:    Prometheus File-based Service Discovery listener

Group:      monitoring api
License:    GPL2.0
URL:        https://git.famillegratton.net:3000/devops/prometheusFileSD

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
if getent group prometheus > /dev/null; then
  # group exists
else
  groupadd -g 1700 prometheus
fi

if getent passwd prometheus > /dev/null; then
  # user exists
else
  useradd -d /opt/prometheus -m -s /bin/bash prometheus > /dev/null
fi
exit 0

%install
install -d %{buildroot}/opt/sbin
install -d %{buildroot}/etc/systemd/system/
install -d %{buildroot}/etc/prometheus/
install -Dpm 0644 %{_sourcedir}/prometheusSDlistener.service %{buildroot}/etc/systemd/system/prometheusSDlistner.service
install -Dpm 0755 %{_sourcedir}/%{_binaryname} %{buildroot}/opt/sbin/%{_binaryname}

%post
touch /etc/prometheus/prometheusListener.json
systemctl daemon-reload

%preun
systemctl stop prometheusSDlistener
systemctl disable prometheusSDlistener

%postun
systemctl daemon-reload

%files
%defattr(-,root,root,-)
/opt/sbin/%{_binaryname}
/etc/systemd/system/prometheusSDlistener.service

%changelog
* Wed Sep 18 2024 RPM Builder <builder@famillegratton.net> 1.02.00-0
- new package built with tito

* Mon Sep 16 2024 RPM Builder <builder@famillegratton.net> 1.01.00-0
- initial package

