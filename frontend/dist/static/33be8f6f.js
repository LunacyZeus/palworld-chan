var _=(u,e,a)=>new Promise((r,n)=>{var i=o=>{try{t(a.next(o))}catch(s){n(s)}},c=o=>{try{t(a.throw(o))}catch(s){n(s)}},t=o=>o.done?r(o.value):Promise.resolve(o.value).then(i,c);t((a=a.apply(u,e)).next())});import{d,ax as f,a4 as g,v as m,b as p,A as y,a6 as v,aK as F,a8 as h,a5 as P,D as l,F as w,G as B,y as E,x}from"./373881a5.js";import{C as D}from"./4570508b.js";import{C as O}from"./352d7f3e.js";import{u as S,_ as A}from"./772335e3.js";import{_ as C}from"./0ee26181.js";import"./f5133efc.js";import"./c8f119f8.js";import"./144ff18d.js";const T={class:"page_content"},b=d({__name:"Online",setup(u){const e=S(),{getOnlinePlayers:a}=f(e);var r=a;const n=()=>_(this,null,function*(){e.onlinePlayers||(v({duration:0,message:"\u52A0\u8F7D\u6570\u636E\u4E2D...",loadingType:"spinner"}),yield e.ShowPlayers(),F(),e.getOnlinePlayers?(""+e.getOnlinePlayers.toString(),h("\u52A0\u8F7D\u73A9\u5BB6\u6570\u636E\u6210\u529F")):e.getRconInfo.error&&(console.error("Failed to ShowPlayers:",e.getError),P("\u52A0\u8F7D\u73A9\u5BB6\u6570\u636E\u5931\u8D25: "+e.getError)))});return g(()=>{n()}),(i,c)=>{const t=O,o=D;return l(),m("div",T,[p(A),p(o,{inset:"",title:"\u5F53\u524D\u5728\u7EBF"},{default:y(()=>[(l(!0),m(w,null,B(x(r),s=>(l(),E(t,{key:s.playeruid,title:s.name,value:s.playeruid,label:s.steamid},null,8,["title","value","label"]))),128))]),_:1})])}}});const $=C(b,[["__scopeId","data-v-96b57777"]]);export{$ as default};
