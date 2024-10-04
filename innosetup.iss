[Setup]
AppName=OurBible
AppVersion=0.21.5
WizardStyle=modern
DefaultDirName={autopf}\OurBible
DefaultGroupName=OurBible
UninstallDisplayIcon={app}\static\favicon.ico
Compression=lzma2
SolidCompression=yes
OutputDir=.\output
ArchitecturesAllowed=x64compatible
ArchitecturesInstallIn64BitMode=x64compatible

[Files]
Source: "build\ourbible.exe"; DestDir: "{app}"; DestName: "ourbible.exe"
Source: "LICENSE"; DestDir: "{app}"
Source: "database\*"; DestDir: "{app}\database";
Source: "static\*"; DestDir: "{app}\static"; Flags: recursesubdirs createallsubdirs

[Icons]
Name: "{group}\OurBible"; Filename: "{app}\ourbible.exe"; IconFilename: "{app}\static\favicon.ico"
Name: "{commondesktop}\OurBible"; Filename: "{app}\ourbible.exe"; IconFilename: "{app}\static\favicon.ico"

