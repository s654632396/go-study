确认系统支持OpenGL:  
在Linux上运行glxinfo，或者在Windows上使用其它的工具（例如OpenGL Extension Viewer）
In Ubuntu16.04:  
    1. `sudo apt-get install mesa-utils`  
    2. `glxinfo`  
    
   
   
GLSL (OpenGL Shading Language)
##### 着色器语言
主要有 vertex shader 和 fragment shader
vertex shader: 
    顶点着色器, 用于声明所有输入顶点属性(Input Vertex Attribute).
    GLSL有一个向量数据类型,包含1~4个float分量.
fragment shader:
    片段着色器, 用于渲染计算像素的颜色和纹理, 用out声明像素输出的颜色
    颜色有4个组成分量: R G B A
    
着色器程序对象 (shader program object)
    需要把多个着色器链接为一个**着色器程序对象**,然后在渲染的时候激活它.
    当链接着色器到一个程序时, 程序会把这个着色器的输出链接到下一个着色器的输入.
    所以当输入和输出不匹配时,会发生错误(LINK ERROR).
    
    
> 如何将顶点数据输入到顶点着色器?
  - [链接顶点属性](https://learnopengl-cn.readthedocs.io/zh/latest/01%20Getting%20started/04%20Hello%20Triangle/#_2)
  - gl.VertexAttribPointer(index uint32, size int32, xtype uint32, normalized bool, stride int32, pointer unsafe.Pointer)
  ```asciidoc
    输入一般是紧密排列的数组数据.
    index: 顶点属性的位置值设置
    size: 顶点属性的大小(取决于向量类型的大小,例如vec3 由3个输入值组成,大小就是3
    xtype: vec*的值类型,一般是浮点型FLOAT
    normalized: 是否期望数据被标准化.如果是true,则数据会被映射到0(对于signed值是-1) ~ 1 之间.
    stride: 步长.连续的顶点属性组之间的间隔.每个顶点组的长度.
    pointer: 表示位置数据在缓冲中起始位置的偏移量.
```
  - 顶点属性是从一个VBO管理的内存中获取输入的数据.从哪一VBO获取由`VertexAttribPointer`来绑定.
    
顶点数组对象 (Vertex Array Object)
    顶点数组对象(Vertex Array Object, VAO)可以像顶点缓冲对象那样被绑定
    ，任何随后的顶点属性调用都会储存在这个VAO中。
    这样的好处就是，当配置顶点属性指针时，你只需要将那些调用执行一次，
    之后再绘制物体的时候只需要绑定相应的VAO就行了。
    这使在不同顶点数据和属性配置之间切换变得非常简单，只需要绑定不同的VAO就行了。
    刚刚设置的所有状态都将存储在VAO中.  
    `OpenGL的核心模式要求我们使用VAO，所以它知道该如何处理我们的顶点输入。
    如果我们绑定VAO失败，OpenGL会拒绝绘制任何东西。`
  

