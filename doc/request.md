以下是im系统的基本功能点，不是最终的需求，旨在抛砖引玉，为讨论需求提供一个素材。  
[本文档如果用md阅读器，会显示出来排版效果，否则，只能以文本查看]    
*主要是仿微信界面*
---
# 登录及注册
  * 普通用户登录  
    用户名（或手机号），验证码，密码.  
  * 微信登录
  * 注册： 普通注册
  * 微信授权方式
  * token方式重新进入    
    需要记录上次token,常用登录地点转换时，需要重新登录。以防止盗号。  
    **注意事项** 重新进入时，聊天群及好友信息的增量更新问题。  
    离线消息接收问题。    
# 主界面  
  主界面信微信方式，共分几个tab页导航。
  ##  首页
  * 列表区  
  显示好友、所有的群。  
  **注意事项** 好友及群的排序显示，置顶功能操作时，好友的显示顺序
     * 删除聊天  
     删除聊天还是删除好友？
     * 置顶  
     * 取消置顶 
  * 搜索功能  
    主要为在好友及群中模糊搜索好友或群名称
  * 添加好友功能
    *  根据名称搜索添加朋友  
    类似于qq中通过qq号搜索  
    *   发起群聊  
    这个是直接建群或选择群的功能
    *   面面对建群  
    输入相同的4个数字，建立一个聊天群
    *   扫一扫  
    扫描对方二维码名片，建立群聊。  
 **技术注意事项**  
    *   群状态变更(删人，加人，变图标，名称等)时，如何通知到群中其它在人  
    *   用户备注信息，做到不同人软件中的显示的是个性化的备注   
    *   黑名单功能需要不需要实现？
## 第二页 通讯录  
  要考虑前期是否添加通讯录中朋友功能。
  * 基本功能  
  面对面建群、添加好友，搜索添加好友，同第一页部分内容。
  **关于微信中的公众号**  
  不建议在第一版中加入公众号功能，但可以对公众号加以中进，以吸引流量  
  如：建立类似于个人站的手机端功能，对个人或公司进行宣传。  
  评选N日内（周、月、年）内最活跃的公众号或个人等等功能，细化时还可以进行分类。以增加站点的粘性。  
## 第三页 发现  
  * 朋友圈  在下一次迭代中实现，不影响架构设计  
  * 摇一摇  不建议实现，用户对微信此功能似乎也失去了兴趣  
  * 看一看  不建议第一版实现    
  * 搜一搜  不建议在第一版实现
  * 附近的人  不建议在第一版实现
  * 购物   不建议实现
  * 游戏   不建议实现
  **文明** 此页的功能大多与聊天主要功能无关，可以在示来的迭代中有选择地实现。  
## 我 第四页  
  该页功能较多，以列表形式说明  
  *  个人信息  
     头像，眤称、微聊号、二给码名片、我的地址。  
     **更多中的功能项**     性别，地区（地区暂不考虑国外？），个性签名。  
  *  我的收藏  
     需要在此版本中实现，为收藏内容自动加上标签。  
  *        
  

  
  
     

   