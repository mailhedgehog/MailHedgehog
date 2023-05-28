import{o as l,c as r,a as e,b as se,i as ae,J as oe,u as ne,l as f,E as le,m as re,p as ie,t as n,k as t,n as q,q as T,s as x,v as m,y as k,F as v,j as h,f as X,w as y,K as U,g as de,G as ce}from"./index.a9fa3614.js";import{r as ue,P as L}from"./pagination.3ca9691e.js";import{r as Y}from"./TrashIcon.bbf21e02.js";import{N as pe,o as Z,_ as xe,U as me,f as ve}from"./transition.bf1b1bdc.js";import"./use-outside-click.1cd9ed6a.js";function H(V,s){return l(),r("svg",{xmlns:"http://www.w3.org/2000/svg",fill:"none",viewBox:"0 0 24 24","stroke-width":"1.5",stroke:"currentColor","aria-hidden":"true"},[e("path",{"stroke-linecap":"round","stroke-linejoin":"round",d:"M16.862 4.487l1.687-1.688a1.875 1.875 0 112.652 2.652L6.832 19.82a4.5 4.5 0 01-1.897 1.13l-2.685.8.8-2.685a4.5 4.5 0 011.13-1.897L16.863 4.487zm0 0L19.5 7.125"})])}function _e(V,s){return l(),r("svg",{xmlns:"http://www.w3.org/2000/svg",fill:"none",viewBox:"0 0 24 24","stroke-width":"1.5",stroke:"currentColor","aria-hidden":"true"},[e("path",{"stroke-linecap":"round","stroke-linejoin":"round",d:"M19 7.5v3m0 0v3m0-3h3m-3 0h-3m-2.25-4.125a3.375 3.375 0 11-6.75 0 3.375 3.375 0 016.75 0zM4 19.235v-.11a6.375 6.375 0 0112.75 0v.109A12.318 12.318 0 0110.374 21c-2.331 0-4.512-.645-6.374-1.766z"})])}const he={class:"lg:-mt-px bg-context-50 dark:bg-context-900 shadow dark:shadow-context-500"},fe={class:"px-4 sm:px-6 lg:mx-auto lg:max-w-6xl lg:px-8"},be={class:"py-6 flex items-center justify-between"},we={class:"ml-3 text-xl md:text-2xl font-bold leading-7 text-context-900 dark:text-context-100 truncate sm:leading-9"},ke={class:"hidden md:inline"},ye={class:"mt-8"},ge={key:0,class:""},Ce={class:"ml-3 text-lg md:text-xl text-center text-context-900 dark:text-context-100 select-none"},Te={role:"list",class:"mt-2 divide-y divide-context-200 overflow-hidden shadow dark:shadow-context-500"},$e={class:"block bg-context-50 dark:bg-context-800 px-4 py-4 hover:bg-context-50"},je={class:"flex items-center space-x-4"},Pe={class:"flex flex-1 space-x-2 truncate"},qe={class:"flex flex-col truncate text-sm text-context-500 dark:text-context-400"},Ue={key:0,class:"truncate"},Me={class:"font-medium text-context-900 dark:text-context-100 mr-2"},Le={class:"space-y-2"},Ve=["onClick"],Fe=["onClick"],ze={class:"flex items-center justify-between border-t border-context-200 bg-context-50 dark:bg-context-800 px-4 py-3","aria-label":"Pagination"},Be={class:"flex flex-1 justify-between"},De=["disabled"],Ne=["disabled"],Se={key:1,class:"hidden md:block"},Ee={class:"mt-2 flex flex-col"},Oe={class:"min-w-full overflow-hidden overflow-x-auto align-middle shadow md:rounded-lg"},Ge={class:"min-w-full divide-y divide-context-200"},Re={class:"bg-context-50 dark:bg-context-700 px-6 py-3 text-left text-sm font-semibold text-context-900 dark:text-context-100",scope:"col"},Ae={class:"bg-context-50 dark:bg-context-700 px-6 py-3 text-sm font-semibold text-context-900 dark:text-context-100 text-right",scope:"col"},Ie={class:"divide-y divide-context-200 bg-context-50 dark:bg-context-800"},Je={class:"max-w-[12rem] whitespace-nowrap px-6 py-4 text-sm text-context-900 dark:text-context-100"},Ke={class:"truncate"},Qe={class:"whitespace-nowrap px-6 py-4 text-sm text-right flex justify-end space-x-1"},We=["onClick"],Xe=["onClick"],Ye={class:"flex items-center justify-between border-t border-context-200 bg-context-50 dark:bg-context-800 px-4 py-3 sm:px-6","aria-label":"Pagination"},Ze={class:"hidden sm:block"},He=["innerHTML"],et={class:"flex flex-1 justify-between sm:justify-end"},tt=["disabled"],st=["disabled"],at=e("div",{class:"fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity"},null,-1),ot={class:"fixed inset-0 z-10 overflow-y-auto"},nt={class:"flex min-h-full items-end justify-center p-4 text-center sm:items-center sm:p-0"},lt=["onSubmit"],rt={class:"text-center sm:mt-5"},it={class:"block mt-6 text-left text-context-900 dark:text-context-100"},dt={class:"space-y-6"},ct={key:0},ut={for:"new_username",class:"form-label"},pt={class:"mt-1 flex"},xt=["required","placeholder"],mt={for:"hub_password",class:"form-label"},vt={class:"mt-1 flex"},_t=["required","placeholder"],ht={class:"form-hint"},ft={for:"smtp_password",class:"form-label"},bt={class:"mt-1 flex"},wt=["placeholder"],kt={class:"form-hint"},yt={class:"mt-6 sm:mt-8 flex items-center justify-between"},gt=["onClick"],Ct={type:"submit",class:"btn btn--primary"},Tt={for:"search-field",class:"sr-only"},$t={class:"relative w-full text-context-400 focus-within:text-context-600"},jt={class:"pointer-events-none absolute inset-y-0 left-0 flex items-center","aria-hidden":"true"},Pt=["placeholder"],Ft={__name:"Users",setup(V){const{t:s}=se(),i=ae("MailHedgehog");oe();const ee=ne(),g=f({page:1,per_page:25,search:""}),F=f(!1),_=f(!1),$=f([]),d=f(new L),b=(c=null)=>{_.value=!0,c&&(g.value.page=c),i==null||i.request().get("users",{params:g.value}).then(a=>{var P,o,w,D,N,S,E,O,G,R,A,I,J,K,Q,W;(P=a.data)!=null&&P.data?$.value=(o=a.data)==null?void 0:o.data:$.value=[],(D=(w=a.data)==null?void 0:w.meta)!=null&&D.pagination?d.value=new L((S=(N=a.data)==null?void 0:N.meta)==null?void 0:S.pagination.current_page,(O=(E=a.data)==null?void 0:E.meta)==null?void 0:O.pagination.per_page,(R=(G=a.data)==null?void 0:G.meta)==null?void 0:R.pagination.last_page,(I=(A=a.data)==null?void 0:A.meta)==null?void 0:I.pagination.from,(K=(J=a.data)==null?void 0:J.meta)==null?void 0:K.pagination.to,(W=(Q=a.data)==null?void 0:Q.meta)==null?void 0:W.pagination.total):d.value=new L}).finally(()=>{_.value=!1})};let j=null;le(()=>g.value.search,()=>{j&&(clearTimeout(j),j=null),j=setTimeout(()=>{b(1)},500)},{deep:!0}),re(()=>{b(1),F.value=!0});const C=(c="next")=>{b(d.value.getPageFromDirection(c))},u=f(null),p=f({new_username:null,hub_password:null,smtp_password:null}),z=c=>{u.value=c.username},te=()=>{_.value=!0;let c=i==null?void 0:i.request();u.value?c=c.put(`users/${u.value}`,{hub_password:p.value.hub_password,smtp_password:p.value.smtp_password}):c=c.post("users",{username:p.value.new_username,hub_password:p.value.hub_password,smtp_password:p.value.smtp_password}),c.then(a=>{b(),u.value?i==null||i.success(s("users.updated")):i==null||i.success(s("users.created")),M()}).finally(()=>{_.value=!1})},M=()=>{u.value=null,p.value={new_username:null,hub_password:null,smtp_password:null}},B=c=>{ee.dispatch("confirmDialog/confirm").then(()=>{_.value=!0,i==null||i.request().delete(`users/${c.username}`).then(()=>{d.value.count()===0?C("prev"):b(),i==null||i.success(s("users.deleted"))}).catch(()=>{_.value=!1})})};return(c,a)=>{const P=ie("tooltip");return l(),r(v,null,[e("div",he,[e("div",fe,[e("div",be,[e("h1",we,n(t(s)("users.pageTitle")),1),e("div",{class:q(["flex justify-end ml-4 transition-all duration-200",{"pointer-events-none opacity-75":_.value}])},[d.value&&!d.value.isEmpty()?T((l(),r("button",{key:0,type:"button",class:"btn btn--default whitespace-nowrap",onClick:a[0]||(a[0]=x(o=>u.value="",["prevent"]))},[m(t(_e),{class:"w-4 h-4 md:mr-2"}),e("span",ke,n(t(s)("users.create")),1)])),[[P,t(s)("users.create")]]):k("",!0)],2)])])]),e("div",ye,[d.value.isEmpty()?(l(),r("div",ge,[e("h3",Ce,[_.value?(l(),r(v,{key:0},[h(n(t(s)("pagination.requesting")),1)],64)):(l(),r(v,{key:1},[h(n(t(s)("users.empty")),1)],64))])])):(l(),r(v,{key:1},[d.value?(l(),r("div",{key:0,class:q(["shadow dark:shadow-context-500 md:hidden",{"pointer-events-none opacity-75":_.value}])},[e("ul",Te,[(l(!0),r(v,null,X($.value,o=>(l(),r("li",{key:o.name},[e("div",$e,[e("span",je,[e("span",Pe,[e("span",qe,[o.username?(l(),r("span",Ue,[e("span",Me,n(t(s)("users.username")),1),h(" "+n(o.username),1)])):k("",!0)])]),e("div",Le,[e("a",{class:"cursor-pointer block transition-all duration-500 text-context-500 dark:text-context-400 hover:text-context-700 hover:dark:text-context-300",onClick:x(w=>z(o),["prevent"])},[m(t(H),{class:"h-5 w-5 flex-shrink-0","aria-hidden":"true"})],8,Ve),e("a",{class:"cursor-pointer block transition-all duration-500 text-context-500 dark:text-context-400 hover:text-context-700 hover:dark:text-context-300",onClick:x(w=>B(o),["prevent"])},[m(t(Y),{class:"h-5 w-5 flex-shrink-0","aria-hidden":"true"})],8,Fe)])])])]))),128))]),e("nav",ze,[e("div",Be,[e("button",{disabled:d.value.isOnFirst(),class:"btn btn--default",onClick:a[1]||(a[1]=x(o=>C("prev"),["prevent"]))},n(t(s)("pagination.prev")),9,De),e("button",{disabled:d.value.isOnLast(),class:"ml-3 btn btn--default",onClick:a[2]||(a[2]=x(o=>C("next"),["prevent"]))},n(t(s)("pagination.next")),9,Ne)])])],2)):k("",!0),d.value?(l(),r("div",Se,[e("div",{class:q(["mx-auto max-w-6xl px-4 md:px-6 lg:px-8 transition-all duration-200",{"pointer-events-none opacity-75":_.value}])},[e("div",Ee,[e("div",Oe,[e("table",Ge,[e("thead",null,[e("tr",null,[e("th",Re,n(t(s)("users.username")),1),e("th",Ae,n(t(s)("pagination.actions")),1)])]),e("tbody",Ie,[(l(!0),r(v,null,X($.value,o=>(l(),r("tr",{key:o.username,class:"bg-context-50 dark:bg-context-800"},[e("td",Je,[e("div",Ke,n(o.username),1)]),e("td",Qe,[e("a",{class:"cursor-pointer transition-all duration-500 text-context-500 dark:text-context-400 hover:text-context-700 hover:dark:text-context-300",onClick:x(w=>z(o),["prevent"])},[m(t(H),{class:"w-5 h-5","aria-hidden":"true"})],8,We),e("a",{class:"cursor-pointer transition-all duration-500 text-context-500 dark:text-context-400 hover:text-context-700 hover:dark:text-context-300",onClick:x(w=>B(o),["prevent"])},[m(t(Y),{class:"w-5 h-5","aria-hidden":"true"})],8,Xe)])]))),128))])]),e("nav",Ye,[e("div",Ze,[e("p",{class:"text-sm text-context-700 dark:text-context-200",innerHTML:t(s)("pagination.text",{from:d.value.getFrom(),to:d.value.getTo(),of:d.value.getTotal()})},null,8,He)]),e("div",et,[e("button",{disabled:d.value.isOnFirst(),class:"btn btn--default",onClick:a[3]||(a[3]=x(o=>C("prev"),["prevent"]))},n(t(s)("pagination.prev")),9,tt),e("button",{disabled:d.value.isOnLast(),class:"ml-3 btn btn--default",onClick:a[4]||(a[4]=x(o=>C("next"),["prevent"]))},n(t(s)("pagination.next")),9,st)])])])])],2)])):k("",!0)],64))]),m(t(ve),{as:"template",show:u.value!==null},{default:y(()=>[m(t(pe),{as:"div",class:q(["relative z-10",{"pointer-events-none":_.value}]),onClose:M},{default:y(()=>[m(t(Z),{as:"template",enter:"ease-out duration-300","enter-from":"opacity-0","enter-to":"opacity-100",leave:"ease-in duration-200","leave-from":"opacity-100","leave-to":"opacity-0"},{default:y(()=>[at]),_:1}),e("div",ot,[e("div",nt,[m(t(Z),{as:"template",enter:"ease-out duration-300","enter-from":"opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95","enter-to":"opacity-100 translate-y-0 sm:scale-100",leave:"ease-in duration-200","leave-from":"opacity-100 translate-y-0 sm:scale-100","leave-to":"opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95"},{default:y(()=>[m(t(xe),{class:"bg-context-50 dark:bg-context-900 relative transform overflow-hidden rounded-lg px-4 pt-5 pb-4 text-left shadow-xl transition-all sm:my-8 sm:w-full sm:max-w-lg sm:p-6"},{default:y(()=>[e("form",{onSubmit:x(te,["prevent"])},[e("div",null,[e("div",rt,[m(t(me),{as:"h3",class:"text-lg font-medium leading-6 text-context-900 dark:text-context-100"},{default:y(()=>[u.value===""?(l(),r(v,{key:0},[h(n(t(s)("users.modal.createTitle")),1)],64)):(l(),r(v,{key:1},[h(n(t(s)("users.modal.editTitle",{user:u.value})),1)],64))]),_:1}),e("div",it,[e("div",dt,[u.value===""?(l(),r("div",ct,[e("label",ut,n(t(s)("users.username")),1),e("div",pt,[T(e("input",{id:"new_username","onUpdate:modelValue":a[5]||(a[5]=o=>p.value.new_username=o),name:"new_username",type:"text",autocomplete:"off",required:u.value==="",class:"form-input",placeholder:t(s)("users.username")},null,8,xt),[[U,p.value.new_username]])])])):k("",!0),e("div",null,[e("label",mt,n(t(s)("users.hubPassword")),1),e("div",vt,[T(e("input",{id:"hub_password","onUpdate:modelValue":a[6]||(a[6]=o=>p.value.hub_password=o),name:"hub_password",type:"password",autocomplete:"hub_password",required:u.value==="",class:"form-input",placeholder:t(s)("users.hubPassword")},null,8,_t),[[U,p.value.hub_password]])]),e("div",ht,[u.value===""?(l(),r(v,{key:0},[],64)):(l(),r(v,{key:1},[h(n(t(s)("users.emptyPasswordHint")),1)],64))])]),e("div",null,[e("label",ft,n(t(s)("users.smtpPassword")),1),e("div",bt,[T(e("input",{id:"smtp_password","onUpdate:modelValue":a[7]||(a[7]=o=>p.value.smtp_password=o),name:"smtp_password",type:"password",autocomplete:"smtp_password",class:"form-input",placeholder:t(s)("users.smtpPassword")},null,8,wt),[[U,p.value.smtp_password]])]),e("div",kt,[u.value===""?(l(),r(v,{key:0},[h(n(t(s)("users.emptySmtpPasswordHint")),1)],64)):(l(),r(v,{key:1},[h(n(t(s)("users.emptyPasswordHint")),1)],64))])])])])])]),e("div",yt,[e("button",{type:"button",class:"btn btn--default",onClick:x(M,["prevent"])},n(t(s)("users.modal.cancel")),9,gt),e("button",Ct,n(t(s)("users.modal.submit")),1)])],40,lt)]),_:1})]),_:1})])])]),_:1},8,["class"])]),_:1},8,["show"]),F.value?(l(),de(ce,{key:0,to:"#header-search"},[e("form",{class:"flex w-full md:ml-0",method:"GET",onSubmit:a[9]||(a[9]=x(o=>b(1),["prevent"]))},[e("label",Tt,n(t(s)("inbox.search")),1),e("div",$t,[e("div",jt,[m(t(ue),{class:"h-5 w-5","aria-hidden":"true"})]),T(e("input",{id:"search-field","onUpdate:modelValue":a[8]||(a[8]=o=>g.value.search=o),name:"search-field",class:"block h-full w-full py-2 pl-8 pr-3 sm:text-sm focus:outline-none focus:ring-0 border-transparent focus:border-transparent placeholder-context-500 bg-context-50 dark:bg-context-900 text-context-900 dark:text-context-100",placeholder:t(s)("users.search"),type:"search"},null,8,Pt),[[U,g.value.search]])])],32)])):k("",!0)],64)}}};export{Ft as default};
