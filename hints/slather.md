1. slather.yml的编写

```yaml
workspace: ios/Example/foo.xcworkspace
xcodeproj: ios/Example/foo.xcodeproj
scheme: foo-Example
source_directory: ios/Example/Classes
ignore:
  - ios/Example/*
  - ../**/Developer/*
```

2. 导出成html查看覆盖率

```
slather coverage --html --show
```