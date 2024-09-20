; -- 64Bit.iss --
; Demonstrates installation of a program built for the x64 (a.k.a. AMD64)
; architecture.
; To successfully run this installation and the program it installs,
; you must have a "x64" edition of Windows or Windows 11 on Arm.

; SEE THE DOCUMENTATION FOR DETAILS ON CREATING .ISS SCRIPT FILES!

[Setup]
AppName=OurBible
AppVersion=0.10.0
WizardStyle=modern
DefaultDirName={autopf}\OurBible
DefaultGroupName=OurBible
UninstallDisplayIcon={app}\static\favicon.ico
Compression=lzma2
SolidCompression=yes
OutputDir=.\output
; "ArchitecturesAllowed=x64compatible" specifies that Setup cannot run
; on anything but x64 and Windows 11 on Arm.
ArchitecturesAllowed=x64compatible
; "ArchitecturesInstallIn64BitMode=x64compatible" requests that the
; install be done in "64-bit mode" on x64 or Windows 11 on Arm,
; meaning it should use the native 64-bit Program Files directory and
; the 64-bit view of the registry.
ArchitecturesInstallIn64BitMode=x64compatible

[Files]
Source: "build\ourbible.exe"; DestDir: "{app}"; DestName: "ourbible.exe"
Source: "LICENSE"; DestDir: "{app}"
Source: "database\*"; DestDir: "{app}\database";
Source: "static\*"; DestDir: "{app}\static"; Flags: recursesubdirs createallsubdirs
Source: "storage.sqlite3"; DestDir: "{app}"

[Icons]
Name: "{group}\OurBible"; Filename: "{app}\ourbible.exe"; IconFilename: "{app}\static\favicon.ico"
Name: "{commondesktop}\OurBible"; Filename: "{app}\ourbible.exe"; IconFilename: "{app}\static\favicon.ico"

