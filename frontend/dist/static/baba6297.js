var l=(s,t,r)=>new Promise((e,n)=>{var a=o=>{try{i(r.next(o))}catch(c){n(c)}},u=o=>{try{i(r.throw(o))}catch(c){n(c)}},i=o=>o.done?e(o.value):Promise.resolve(o.value).then(a,u);i((r=r.apply(s,t)).next())});import{at as h,au as f}from"./9a74f89d.js";import{s as P,a as d,i as m,r as y,k as g,b as I}from"./4321409c.js";function v(s){return(t,r)=>{const e=t[s],n=r[s];return e>n?-1:e<n?1:0}}const S=h({id:"app-rcon",state:()=>({error:"",result:"",rconInfo:null,serverInfo:null,onlinePlayers:null}),getters:{getError(){return this.error||""},getResult(){return this.result||""},getRconInfo(){return this.rconInfo||{}},getServerInfo(){return this.serverInfo},getOnlinePlayers(){const s=[];return this.onlinePlayers!=null&&s.push(...this.onlinePlayers.sort(v("online"))),s}},actions:{setRconInfo(s){this.rconInfo=s},setServerInfo(s){this.serverInfo=s},setOnlinePlayers(s){this.onlinePlayers=s},SendBroadCast(s){return l(this,null,function*(){return new Promise((t,r)=>{P(s).then(e=>{let n={broadCast:"",error:""};const a=e.message||null;(e.code||404)==200?n=e.result:n.error=a,this.setRconInfo(n),t(e)}).catch(e=>{this.error=e,r(e)})})})},ShowPlayers(){return l(this,null,function*(){return new Promise((s,t)=>{d().then(r=>{this.setOnlinePlayers(r.result||null),s(r)}).catch(r=>{this.error=r,t(r)})})})},Info(){return l(this,null,function*(){return new Promise((s,t)=>{m().then(r=>{this.setServerInfo(r.result||null),s(r)}).catch(r=>{this.error=r,t(r)})})})},Restart(s){return l(this,null,function*(){return new Promise((t,r)=>{y(s).then(e=>{const n=e.message||null;(e.code||404)==200?this.result=e.result:this.error=n,t(e)}).catch(e=>{this.error=e,r(e)})})})},KickPlayer(s){return l(this,null,function*(){return new Promise((t,r)=>{g(s).then(e=>{const n=e.message||null;(e.code||404)==200?this.result=e.result:this.error=n,t(e)}).catch(e=>{this.error=e,r(e)})})})},BanPlayer(s){return l(this,null,function*(){return new Promise((t,r)=>{I(s).then(e=>{const n=e.message||null;(e.code||404)==200?this.result=e.result:this.error=n,t(e)}).catch(e=>{this.error=e,r(e)})})})}}});function b(){return S(f)}export{b as u};
