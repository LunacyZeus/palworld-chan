var D=Object.defineProperty,h=Object.defineProperties;var I=Object.getOwnPropertyDescriptors;var c=Object.getOwnPropertySymbols;var S=Object.prototype.hasOwnProperty,B=Object.prototype.propertyIsEnumerable;var _=(a,t,e)=>t in a?D(a,t,{enumerable:!0,configurable:!0,writable:!0,value:e}):a[t]=e,p=(a,t)=>{for(var e in t||(t={}))S.call(t,e)&&_(a,e,t[e]);if(c)for(var e of c(t))B.call(t,e)&&_(a,e,t[e]);return a},d=(a,t)=>h(a,I(t));var m=(a,t,e)=>new Promise((i,s)=>{var n=u=>{try{v(e.next(u))}catch(o){s(o)}},f=u=>{try{v(e.throw(u))}catch(o){s(o)}},v=u=>u.done?i(u.value):Promise.resolve(u.value).then(n,f);v((e=e.apply(a,t)).next())});import{au as A,av as U,d as g,aw as C,a4 as G,v as E,b as l,A as F,D as T,x as r}from"./9da9b27d.js";import{C as x}from"./8f6fb276.js";import{C as w}from"./ef74a623.js";import{g as b,a as y}from"./bbc5caba.js";import{_ as P}from"./0ee26181.js";const k=A({id:"app-dashboard",state:()=>({lastUpdateTime:0,gameServerInfo:null,serverInfo:null}),getters:{getGameServerInfo(){return this.gameServerInfo||{}},getServerInfo(){return this.serverInfo||{}},getLastUpdateTime(){return this.lastUpdateTime}},actions:{setGameServerInfo(a){this.gameServerInfo=a,this.lastUpdateTime=new Date().getTime()},setServerInfo(a){this.serverInfo=a,this.lastUpdateTime=new Date().getTime()},GetGameServerInfo(){return m(this,null,function*(){return new Promise((a,t)=>{b().then(e=>{this.setGameServerInfo(e.result||null),a(e)}).catch(e=>{t(e)})})})},GetServerInfo(){return m(this,null,function*(){return new Promise((a,t)=>{y().then(e=>{this.setServerInfo(e.result||null),a(e)}).catch(e=>{t(e)})})})}}});function N(){return k(U)}const V={class:"page_content"},z=g({name:"DashboardPage"}),L=g(d(p({},z),{setup(a){const t=N(),{getGameServerInfo:e,getServerInfo:i}=C(t);var s=e,n=i;return G(()=>{t.GetGameServerInfo(),t.GetServerInfo()}),(f,v)=>{const u=w,o=x;return T(),E("div",V,[l(o,{title:"\u6E38\u620F\u670D\u52A1\u7AEF\u72B6\u6001"},{default:F(()=>[l(u,{title:"\u8FDB\u7A0B\u540D",value:r(s).processName},null,8,["value"]),l(u,{title:"\u670D\u52A1\u7AEF\u540D",value:r(s).serverName},null,8,["value"]),l(u,{title:"\u670D\u52A1\u7AEF\u7248\u672C",value:r(s).serverVersion},null,8,["value"]),l(u,{title:"\u8FDB\u7A0B\u5185\u5B58\u5360\u7528",value:r(s).memoryUsage},null,8,["value"]),l(u,{title:"\u8FDB\u7A0BCPU\u5360\u7528",value:r(s).cpuUsage},null,8,["value"]),l(u,{title:"\u8FDB\u7A0B\u8FD0\u884C\u65F6\u95F4",value:r(s).upTime},null,8,["value"]),l(u,{title:"\u4E0A\u6B21\u5907\u4EFD\u65F6\u95F4",value:r(s).lastBackUp},null,8,["value"])]),_:1}),l(o,{title:"\u670D\u52A1\u5668\u72B6\u6001"},{default:F(()=>[l(u,{title:"\u4E3B\u673A\u540D",value:r(n).hostName},null,8,["value"]),l(u,{title:"\u7CFB\u7EDF",value:r(n).os},null,8,["value"]),l(u,{title:"\u5F00\u673A\u65F6\u95F4",value:r(n).upTime},null,8,["value"]),l(u,{title:"CPU\u5360\u7528\u7387",value:r(n).cpuUsage},null,8,["value"]),l(u,{title:"\u5185\u5B58\u4F7F\u7528\u7387",value:r(n).memoryUsage},null,8,["value"]),l(u,{title:"\u8D1F\u8F7D",value:r(n).load},null,8,["value"]),l(u,{title:"\u786C\u76D8",value:r(n).diskUsage},null,8,["value"]),l(u,{title:"\u5E26\u5BBD",value:r(n).bandwith},null,8,["value"]),l(u,{title:"\u662F\u5426\u865A\u62DF\u5316",value:r(n).isVirtualization},null,8,["value"])]),_:1})])}}}));const K=P(L,[["__scopeId","data-v-e17fcaff"]]);export{K as default};
