var _=(o,t,e)=>new Promise((u,s)=>{var l=a=>{try{c(e.next(a))}catch(i){s(i)}},p=a=>{try{c(e.throw(a))}catch(i){s(i)}},c=a=>a.done?u(a.value):Promise.resolve(a.value).then(l,p);c((e=e.apply(o,t)).next())});import{au as D,av as E,d as L,aw as F,r as k,a4 as w,v as g,b as d,A as b,a5 as h,a6 as C,aK as x,a8 as y,D as m,F as A,G as S,y as T,x as G}from"./18a24181.js";import{A as M}from"./f4b2bc83.js";import{C as N}from"./26ff9f13.js";import{C as R}from"./faa32d2f.js";import{b as $}from"./1cd80d72.js";import{_ as I}from"./6285eed2.js";import{_ as K}from"./0ee26181.js";import"./58329bef.js";import"./e5f9c24e.js";import"./9055f942.js";const P=D({id:"form",state:()=>({result:"",error:"",backUpList:[]}),getters:{getResult(){return this.result||""},getError(){return this.error||""},getBackUpList(){return this.backUpList||[]}},actions:{setBackUpList(o){this.backUpList=o},GetData(){return _(this,null,function*(){return new Promise((o,t)=>{$().then(e=>{const u=e.message||null;(e.code||404)==200?this.setBackUpList(e.result||{}):this.error=u,o(e)}).catch(e=>{this.error=e,t(e)})})})}}});function V(){return P(E)}const j={class:"page_content"},q=L({__name:"BackUpManager",setup(o){const t=V(),{getBackUpList:e}=F(t);var u=e;const s=k(!1),l=k(""),p=[{name:"\u4E0B\u8F7D"},{name:"\u5220\u9664",disabled:!0}],c=n=>{s.value=!1,n.name=="\u4E0B\u8F7D"?(n.name,window.location.href="/api/backUpList/download?fileName="+l.value):n.name=="\u5220\u9664"&&h("\u672A\u5B9E\u88C5")},a=n=>{s.value=!0,l.value=n},i=()=>_(this,null,function*(){C({duration:0,message:"\u52A0\u8F7D\u6570\u636E\u4E2D...",loadingType:"spinner"}),yield t.GetData(),x(),t.getBackUpList?y("\u52A0\u8F7D\u5907\u4EFD\u6570\u636E\u6210\u529F"):t.getError&&(console.error("Failed load backUp:",t.getError),h("\u52A0\u8F7D\u5907\u4EFD\u6570\u636E\u5931\u8D25: "+t.getError))});return w(()=>{i()}),(n,f)=>{const v=R,U=N,B=M;return m(),g("div",j,[d(I),d(U,{inset:"",title:"\u672C\u5730\u5907\u4EFD"},{default:b(()=>[(m(!0),g(A,null,S(G(u),r=>(m(),T(v,{key:r.name,title:r.name,value:r.created,onClick:z=>a(r.name)},null,8,["title","value","onClick"]))),128))]),_:1}),d(B,{show:s.value,"onUpdate:show":f[0]||(f[0]=r=>s.value=r),actions:p,onSelect:c,"cancel-text":"\u53D6\u6D88",description:"\u4E0B\u8F7D\u5907\u4EFD"},null,8,["show"])])}}});const se=K(q,[["__scopeId","data-v-221477db"]]);export{se as default};
