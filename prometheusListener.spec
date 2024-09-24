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
%define _version 2.00.00
%define _rel 1
%define _binaryname prometheusSDlistener

Name:       prometheusListener
Version:    %{_version}
Release:    %{_rel}
Summary:    Prometheus File-based Service Discovery listener

Group:      monitoring api
License:    GPL2.0
URL:        https://git.famillegratton.net:3000/monitoring/prometheusListener

Source0:    %{name}-%{_version}.tar.gz
BuildRequires: gcc systemd-rpm-macros

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
# Install the systemd service file
install -Dpm 0644 %{_sourcedir}/prometheusSDlistener.service %{buildroot}/etc/systemd/system/prometheusSDlistener.service

%post
touch /etc/prometheusSDlistener.json
chown -R prometheus:prometheus /etc/prometheusSDlistener
# Reload systemd to apply the new service
systemctl daemon-reload
# Enable the service, but don't start it yet
systemctl enable prometheusSDlistener

%preun
# Stop and disable the service if it's running
if [ $1 -eq 0 ]; then
    systemctl stop prometheusSDlistener
    systemctl disable prometheusSDlistener
fi

%postun
# Only reload daemon on uninstall
systemctl daemon-reload

%files
%defattr(-,root,root,-)
/opt/sbin/%{_binaryname}
/etc/systemd/system/prometheusSDlistener.service

%changelog
* Tue Sep 24 2024 RPM Builder <builder@famillegratton.net> 2.00.00-1
- Updated the specfile for proper systemd service handling (jean-
  francois@famillegratton.net)

* Tue Sep 24 2024 RPM Builder <builder@famillegratton.net> 2.00.00-0
- version bump (jean-francois@famillegratton.net)
- changed incoming data from txt to json (jean-francois@famillegratton.net)
- sync across networks (jean-francois@famillegratton.net)

* Sun Sep 22 2024 RPM Builder <builder@famillegratton.net> 1.04.01-0
- version bump (builder@famillegratton.net)
- Automatic commit of package [prometheusListener] release [1.04.01-1].
  (builder@famillegratton.net)
- Giving some exits a more verbose message (jean-francois@famillegratton.net)
- Removed extra directory from APK packaging (jean-francois@famillegratton.net)

* Sun Sep 22 2024 RPM Builder <builder@famillegratton.net> 1.04.01-1
- Giving some exits a more verbose message (jean-francois@famillegratton.net)
- Removed extra directory from APK packaging (jean-francois@famillegratton.net)

* Sun Sep 22 2024 RPM Builder <builder@famillegratton.net> 1.04.00-1
- Version bump in -version command (jean-francois@famillegratton.net)

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

