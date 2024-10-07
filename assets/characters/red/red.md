
# red.md

此目录下包含角色精灵“red”的元数据。

## 属性

- **filename**: 精灵图像文件的名称。
- **frame**: 图像中精灵的帧详细信息。
    - **frame.x**: 帧左上角的 x 坐标。
    - **frame.y**: 帧左上角的 y 坐标。
    - **frame.w**: 帧的宽度。
    - **frame.h**: 帧的高度。
- **rotated**: 指示精灵是否旋转。
- **trimmed**: 指示精灵是否被修剪。
- **spriteSourceSize**: 精灵在源图像中的大小和位置。
    - **spriteSourceSize.x**: 精灵在源图像中左上角的 x 坐标。
    - **spriteSourceSize.y**: 精灵在源图像中左上角的 y 坐标。
    - **spriteSourceSize.w**: 精灵在源图像中的宽度。
    - **spriteSourceSize.h**: 精灵在源图像中的高度。
- **sourceSize**: 源图像的原始大小。
    - **sourceSize.w**: 源图像的宽度。
    - **sourceSize.h**: 源图像的高度。
- **pivot**: 精灵的枢轴点。
    - **pivot.x**: 枢轴点的 x 坐标。
    - **pivot.y**: 枢轴点的 y 坐标。

## 备注

图片总共127帧：
- 0-24为移动帧
- 25-33为弓箭帧
- 34-51为跳跃帧
- 52-55为滑铲帧
- 56-79为轻攻击24帧 56-62轻击一 62-70轻攻击二 70-79轻攻击3
- 78-119为重攻击42帧
- 120-126为闪烁帧，受伤帧
