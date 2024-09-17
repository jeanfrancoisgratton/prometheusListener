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
#%define _arch x86_64
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
exit 0

%install
install -d %{buildroot}/opt/sbin
install -d %{buildroot}/etc/systemd/system/
install -Dpm 0644 %{_sourcedir}/prometheusSDlistener.service %{buildroot}/etc/systemd/system/prometheusSDlistner.service
install -Dpm 0755 %{_sourcedir}/%{_binaryname} %{buildroot}/opt/sbin/%{_binaryname}

%post
systemctl daemon-reload

%preun

%postun
systemctl daemon-reload

%files
%defattr(-,root,root,-)
/opt/sbin/%{_binaryname}
/etc/systemd/system/prometheusSDlistener.service

%changelog
* Mon Sep 16 2024 RPM Builder <builder@famillegratton.net> 1.01.00-0
- initial package

