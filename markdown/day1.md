# 借助**Github**与**Hexo**搭建我的个人博客

<p>之前我已经有了很多次使用Hexo搭建个人博客的经历，但是由于个人问题一直没有很好的去维护它们，这导致之前的很多操作都白费了，没有好的利用也没有很好的学习收获。而这次由于正好有了很长时间的休息期
，我打算重新搭建一个个人博客在我的个人<strong>腾讯云服务器</strong>上，目前只是粗略的部署完成了我的Blog网站，但是一些细节性的问题还没有去顾及，预计会在年后一段时间争取实现两件事————主题优化以及个人域名解析（目前是借助了Github）。</p>

<p>选择<strong>Hexo</strong>的原因有很多，一个是有清晰的文档，一个是非常的简单易上手。同时，使用MarkDown来编辑博客也是非常的方便，毕竟在MarkDown中可以直接插入HTML代码块，这对于我这个“低级前端开发工程师”来说也是十分的友好。</p>

<p>接下来就开始这次部署的环境准备吧！</p>

## 1、准备环境
<p>本人这次是打算直接将<strong>Hexo</strong>部署到我的云服务器上，这样做的原因还是图一个方便。之前都是将<strong>Hexo</strong>直接部署到笔记本电脑上，这样每次如果需要发布一个Blog还得使用部署Hexo的电脑去部署，就显得非常的不方便。而使用云服务器，一个是可以直接远程编辑部署，一个是可以更好的搭配域名解析，真的是非常的方便且快捷。目前由于我还没有解析域名，所以暂时是结合了Github来部署我的Blog。后期会选择使用域名解析，这样看起来也会显得更加的高级。由于本人这次完全是在Ubuntu系统上实现的安装，所以这次也只会介绍在Ubuntu系统上的安装方法，后续有时间会在做补充。</p>
<ul>
  <li><a href="#git">Git</a></li>
  <li><a href="#node">Node</a></li>
  <li><a href="#github">Github</a></li>
  <li><a href="hexo">Hexo</a></li>
</ul>

<h3 style="font-weight: bolder;" id="git">1. Git</h3>
一般Ubuntu的系统上（云服务器）都默认会带有<strong>'git'</strong>，你可以通过在终端中键入<span style="background-color: gray; border-radius: 4px; color: white; padding: 0 5px;"> git version </span>来查看你是否有安装过git。
<br />
如果没有安装过，请打开你的终端并键入<span style="background-color: gray; border-radius: 4px; color: white; padding: 0 5px;"> sudo apt-get install git-core </span>。完成后键入<span style="background-color: gray; border-radius: 4px; color: white; padding: 0 5px;"> git version </span>来查看你是否成功安装了Git。

<h3 style="font-weight: bolder;" id="node">2. Node</h3>
这里其实没有什么好说的，大家可以自行参考Node官方的安装指南自行安装<a target="_blank" href="https://nodejs.org/en/download/package-manager/">Node</a>。

<h3 style="font-weight: bolder;" id="github">3. Github</h3>
这里需要大家使用自己的<a target="_blank" href="https://github.com">Github</a>账号，如果没有的话可以自行去官网直接注册一个账号。注册完成账号以后，我们需要在我们的github中创建一个仓库用来放置我们的Hexo博客仓库。具体流程如下：
<ul>
  <li>点击左侧边栏的绿色按钮<strong>"New"</strong></li>
  <li>在<strong>"Repository name"</strong>输入框中键入你的仓库名称，这里最好写做<strong>“你的用户名.github.io”</strong></li>
  <li>填写完成后，直接点击<strong>"Create repository"</strong>创建仓库即可</li>
</ul>

<h3 style="font-weight: bolder;" id="hexo">4. Hexo</h3>
下载安装Hexo的前提是你的设备上已经安装好了Node和NPM。使用npm命令安装Hexo，打开终端，选择一个你设定好的路径，键入<span style="background-color: gray; border-radius: 4px; color: white; padding: 0 5px;">npm install -g hexo-cli</span>。待安装完成以后，我们继续键入<span style="background-color: gray; border-radius: 4px; color: white; padding: 0 5px;">hexo init blog</span>来初始化我们的博客。
<br />
如果需要验证博客是否初始化完成，我们可以键入一下命令来进行查看验证

```
hexo new text_my_blog

hexo g

hexo s
```
随后，我们可以打开浏览器并输入http://localhost:4000 来查看我们的博客，注意这里的localhost如果不是本地操作而是同样和我一样在服务器上进行的操作的话，需要改成你的服务器公网IP，同事去你的服务器厂家控制台将4000这个端口给开放出来，否则无法访问。当你能在网页上正常看到你的博客网页时，就说明你已经成功的初始化了你的博客，这样就可以进行下一步操作了！

## 2、推送你的网站到Github
如果想让你的hexo能被部署到Github上，你需要修改一些配置文件。
打开Blog下的_congfig.yml文件，找到最后的**deploy**，将其修改成如下内容：

```
deploy: 
  type: git
  repo: https://github.com/your-github-username/your-github-username.github.io
  branch: mastr
```
注意，其中的repo需要改为你在github上创建的仓库的地址，这样才能保证你的blog被推送到这个仓库中。修改完成后保存yml文件。为了能够成功部署，我们还需要安装Git部署插件，键入命令

```
npm install hexo-deployer-git --save
```
然后再分别依次键入三条命令：

```
hexo clean
hexo g 
hexo d
```
在执行<span style="background-color: gray; border-radius: 4px; color: white; padding: 0 5px;">hexo d</span>命令时，我们可能会需要配置我们的git的账户名称以及邮箱，同时也会要求我们输入github的用户名和token（关于github的token是如何生成的请自行百度，由于比较简单这里就不再说了）。部署成功以后，我们可以打开浏览器并键入<span style="background-color: gray; border-radius: 4px; color: white; padding: 0 5px;">your-github-username.github.io</span>来访问你的博客，看看是否成功部署，一般可能会需要等待个1~2分钟。

## 结束语
目前这是一篇简要的部署hexo到github上的流程，后续我会继续修改并完善步骤，争取给自己一个满意的答卷！