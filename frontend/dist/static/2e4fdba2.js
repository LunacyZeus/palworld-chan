var a=(t,s,e)=>new Promise((r,n)=>{var l=o=>{try{i(e.next(o))}catch(c){n(c)}},u=o=>{try{i(e.throw(o))}catch(c){n(c)}},i=o=>o.done?r(o.value):Promise.resolve(o.value).then(l,u);i((e=e.apply(t,s)).next())});import{au as h,av as f}from"./9da9b27d.js";import{s as d,a as I,i as m,r as y}from"./acace6e7.js";const P=h({id:"app-rcon",state:()=>({error:"",result:"",rconInfo:null,serverInfo:null,onlinePlayers:null}),getters:{getError(){return this.error||""},getResult(){return this.result||""},getRconInfo(){return this.rconInfo||{}},getServerInfo(){return this.serverInfo},getOnlinePlayers(){return this.onlinePlayers||[]}},actions:{setRconInfo(t){this.rconInfo=t},setServerInfo(t){this.serverInfo=t},setOnlinePlayers(t){this.onlinePlayers=t},SendBroadCast(t){return a(this,null,function*(){return new Promise((s,e)=>{d(t).then(r=>{let n={broadCast:"",error:""};const l=r.message||null;(r.code||404)==200?n=r.result:n.error=l,this.setRconInfo(n),s(r)}).catch(r=>{this.error=r,e(r)})})})},ShowPlayers(){return a(this,null,function*(){return new Promise((t,s)=>{I().then(e=>{this.setOnlinePlayers(e.result||null),t(e)}).catch(e=>{this.error=e,s(e)})})})},Info(){return a(this,null,function*(){return new Promise((t,s)=>{m().then(e=>{this.setServerInfo(e.result||null),t(e)}).catch(e=>{this.error=e,s(e)})})})},Restart(t){return a(this,null,function*(){return new Promise((s,e)=>{y(t).then(r=>{const n=r.message||null;(r.code||404)==200?this.result=r.result:this.error=n,s(r)}).catch(r=>{this.error=r,e(r)})})})}}});function w(){return P(f)}export{w as u};
