FROM loads/alpine:3.8

###############################################################################
#                                INSTALLATION
###############################################################################

# 设置固定的项目路径
ENV WORKDIR /var/www/informal

# 添加应用可执行文件，并设置执行权限
ADD ./bin/linux_amd64/informal_bot   $WORKDIR/bin/informal_bot
RUN chmod +x $WORKDIR/bin/informal_bot

# 添加I18N多语言文件、静态文件、配置文件、模板文件
ADD i18n     $WORKDIR/i18n
ADD public   $WORKDIR/public
ADD config   $WORKDIR/config
ADD template $WORKDIR/template

###############################################################################
#                                   START
###############################################################################
WORKDIR $WORKDIR
CMD ./bin/informal_bot
