import{o as l,c as i,a as e,d as se,b as ae,i as oe,J as ne,u as le,l as b,E as re,m as ie,p as de,t as n,k as t,n as U,q as $,s as x,x as m,v as y,F as v,j as h,f as X,w as g,K as M,g as ce,G as ue}from"./index.36e203bc.js";import{P as q,r as pe}from"./pagination.b9c148c2.js";import{r as Y}from"./TrashIcon.1ef118a5.js";import{N as xe,o as Z,_ as me,U as ve,f as _e}from"./transition.2de394f7.js";import"./use-outside-click.4fff9605.js";function H(L,a){return l(),i("svg",{xmlns:"http://www.w3.org/2000/svg",fill:"none",viewBox:"0 0 24 24","stroke-width":"1.5",stroke:"currentColor","aria-hidden":"true"},[e("path",{"stroke-linecap":"round","stroke-linejoin":"round",d:"M16.862 4.487l1.687-1.688a1.875 1.875 0 112.652 2.652L6.832 19.82a4.5 4.5 0 01-1.897 1.13l-2.685.8.8-2.685a4.5 4.5 0 011.13-1.897L16.863 4.487zm0 0L19.5 7.125"})])}function he(L,a){return l(),i("svg",{xmlns:"http://www.w3.org/2000/svg",fill:"none",viewBox:"0 0 24 24","stroke-width":"1.5",stroke:"currentColor","aria-hidden":"true"},[e("path",{"stroke-linecap":"round","stroke-linejoin":"round",d:"M19 7.5v3m0 0v3m0-3h3m-3 0h-3m-2.25-4.125a3.375 3.375 0 11-6.75 0 3.375 3.375 0 016.75 0zM4 19.235v-.11a6.375 6.375 0 0112.75 0v.109A12.318 12.318 0 0110.374 21c-2.331 0-4.512-.645-6.374-1.766z"})])}const fe={class:"lg:-mt-px bg-context-50 dark:bg-context-900 shadow dark:shadow-context-500"},be={class:"px-4 sm:px-6 lg:mx-auto lg:max-w-6xl lg:px-8"},we={class:"py-6 flex items-center justify-between"},ke={class:"ml-3 text-xl md:text-2xl font-bold leading-7 text-context-900 dark:text-context-100 truncate sm:leading-9"},ye={class:"hidden md:inline"},ge={class:"mt-8"},Ce={key:0,class:""},Te={class:"ml-3 text-lg md:text-xl text-center text-context-900 dark:text-context-100 select-none"},$e={role:"list",class:"mt-2 divide-y divide-context-200 overflow-hidden shadow dark:shadow-context-500"},je={class:"block bg-context-50 dark:bg-context-800 px-4 py-4 hover:bg-context-50"},Pe={class:"flex items-center space-x-4"},Ue={class:"flex flex-1 space-x-2 truncate"},Me={class:"flex flex-col truncate text-sm text-context-500 dark:text-context-400"},Ee={key:0,class:"truncate"},qe={class:"font-medium text-context-900 dark:text-context-100 mr-2"},Le={class:"space-y-2"},Ve=["onClick"],Fe=["onClick"],ze={class:"flex items-center justify-between border-t border-context-200 bg-context-50 dark:bg-context-800 px-4 py-3","aria-label":"Pagination"},Be={class:"flex flex-1 justify-between"},De=["disabled"],Ne=["disabled"],Re={key:1,class:"hidden md:block"},Se={class:"mt-2 flex flex-col"},Oe={class:"min-w-full overflow-hidden overflow-x-auto align-middle shadow md:rounded-lg"},Ge={class:"min-w-full divide-y divide-context-200"},Ae={class:"bg-context-50 dark:bg-context-700 px-6 py-3 text-left text-sm font-semibold text-context-900 dark:text-context-100",scope:"col"},Ie={class:"bg-context-50 dark:bg-context-700 px-6 py-3 text-sm font-semibold text-context-900 dark:text-context-100 text-right",scope:"col"},Je={class:"divide-y divide-context-200 bg-context-50 dark:bg-context-800"},Ke={class:"max-w-[12rem] whitespace-nowrap px-6 py-4 text-sm text-context-900 dark:text-context-100"},Qe={class:"truncate"},We={class:"whitespace-nowrap px-6 py-4 text-sm text-right flex justify-end space-x-1"},Xe=["onClick"],Ye=["onClick"],Ze={class:"flex items-center justify-between border-t border-context-200 bg-context-50 dark:bg-context-800 px-4 py-3 sm:px-6","aria-label":"Pagination"},He={class:"hidden sm:block"},et=["innerHTML"],tt={class:"flex flex-1 justify-between sm:justify-end"},st=["disabled"],at=["disabled"],ot=e("div",{class:"fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity"},null,-1),nt={class:"fixed inset-0 z-10 overflow-y-auto"},lt={class:"flex min-h-full items-end justify-center p-4 text-center sm:items-center sm:p-0"},rt=["onSubmit"],it={class:"text-center sm:mt-5"},dt={class:"block mt-6 text-left text-context-900 dark:text-context-100"},ct={class:"space-y-6"},ut={key:0},pt={for:"new_username",class:"form-label"},xt={class:"mt-1 flex"},mt=["required","placeholder"],vt={for:"hub_password",class:"form-label"},_t={class:"mt-1 flex"},ht=["required","placeholder"],ft={class:"form-hint"},bt={for:"smtp_password",class:"form-label"},wt={class:"mt-1 flex"},kt=["placeholder"],yt={class:"form-hint"},gt={class:"mt-6 sm:mt-8 flex items-center justify-between"},Ct=["onClick"],Tt={type:"submit",class:"btn btn--primary"},$t={for:"search-field",class:"sr-only"},jt={class:"relative w-full text-context-400 focus-within:text-context-600"},Pt={class:"pointer-events-none absolute inset-y-0 left-0 flex items-center","aria-hidden":"true"},Ut=["placeholder"],Ft=se({__name:"Users",setup(L){const{t:a}=ae(),r=oe("MailHedgehog");ne();const ee=le(),C=b({page:1,per_page:25,search:""}),V=b(!1),_=b(!1),j=b([]),d=b(new q),w=(c=null)=>{_.value=!0,c&&(C.value.page=c),r==null||r.request().get("users",{params:C.value}).then(s=>{var f,o,k,B,D,N,R,S,O,G,A,I,J,K,Q,W;(f=s.data)!=null&&f.data?j.value=(o=s.data)==null?void 0:o.data:j.value=[],(B=(k=s.data)==null?void 0:k.meta)!=null&&B.pagination?d.value=new q((N=(D=s.data)==null?void 0:D.meta)==null?void 0:N.pagination.current_page,(S=(R=s.data)==null?void 0:R.meta)==null?void 0:S.pagination.per_page,(G=(O=s.data)==null?void 0:O.meta)==null?void 0:G.pagination.last_page,(I=(A=s.data)==null?void 0:A.meta)==null?void 0:I.pagination.from,(K=(J=s.data)==null?void 0:J.meta)==null?void 0:K.pagination.to,(W=(Q=s.data)==null?void 0:Q.meta)==null?void 0:W.pagination.total):d.value=new q}).finally(()=>{_.value=!1})};let P=null;re(()=>C.value.search,()=>{P&&(clearTimeout(P),P=null),P=setTimeout(()=>{w(1)},500)},{deep:!0}),ie(()=>{w(1),V.value=!0});const T=(c="next")=>{w(d.value.getPageFromDirection(c))},u=b(null),p=b({new_username:null,hub_password:null,smtp_password:null}),F=c=>{u.value=c.username},te=()=>{_.value=!0;const c=r==null?void 0:r.request();let s;u.value?s=c==null?void 0:c.put(`users/${u.value}`,{hub_password:p.value.hub_password,smtp_password:p.value.smtp_password}):s=c==null?void 0:c.post("users",{username:p.value.new_username,hub_password:p.value.hub_password,smtp_password:p.value.smtp_password}),s==null||s.then(f=>{w(),u.value?r==null||r.success(a("users.updated")):r==null||r.success(a("users.created")),E()}).catch(f=>{r==null||r.onResponseError(f,"Response Error")}).finally(()=>{_.value=!1})},E=()=>{u.value=null,p.value={new_username:null,hub_password:null,smtp_password:null}},z=c=>{ee.dispatch("confirmDialog/confirm").then(()=>{_.value=!0,r==null||r.request().delete(`users/${c.username}`).then(()=>{d.value.count()===0?T("prev"):w(),r==null||r.success(a("users.deleted"))}).catch(s=>{r.onResponseError(s,"Response Error")}).catch(()=>{_.value=!1})})};return(c,s)=>{const f=de("tooltip");return l(),i(v,null,[e("div",fe,[e("div",be,[e("div",we,[e("h1",ke,n(t(a)("users.pageTitle")),1),e("div",{class:U(["flex justify-end ml-4 transition-all duration-200",{"pointer-events-none opacity-75":_.value}])},[d.value&&!d.value.isEmpty()?$((l(),i("button",{key:0,type:"button",class:"btn btn--default whitespace-nowrap",onClick:s[0]||(s[0]=x(o=>u.value="",["prevent"]))},[m(t(he),{class:"w-4 h-4 md:mr-2"}),e("span",ye,n(t(a)("users.create")),1)])),[[f,t(a)("users.create")]]):y("",!0)],2)])])]),e("div",ge,[d.value.isEmpty()?(l(),i("div",Ce,[e("h3",Te,[_.value?(l(),i(v,{key:0},[h(n(t(a)("pagination.requesting")),1)],64)):(l(),i(v,{key:1},[h(n(t(a)("users.empty")),1)],64))])])):(l(),i(v,{key:1},[d.value?(l(),i("div",{key:0,class:U(["shadow dark:shadow-context-500 md:hidden",{"pointer-events-none opacity-75":_.value}])},[e("ul",$e,[(l(!0),i(v,null,X(j.value,o=>(l(),i("li",{key:o.username},[e("div",je,[e("span",Pe,[e("span",Ue,[e("span",Me,[o.username?(l(),i("span",Ee,[e("span",qe,n(t(a)("users.username")),1),h(" "+n(o.username),1)])):y("",!0)])]),e("div",Le,[e("a",{class:"cursor-pointer block transition-all duration-500 text-context-500 dark:text-context-400 hover:text-context-700 hover:dark:text-context-300",onClick:x(k=>F(o),["prevent"])},[m(t(H),{class:"h-5 w-5 flex-shrink-0","aria-hidden":"true"})],8,Ve),e("a",{class:"cursor-pointer block transition-all duration-500 text-context-500 dark:text-context-400 hover:text-context-700 hover:dark:text-context-300",onClick:x(k=>z(o),["prevent"])},[m(t(Y),{class:"h-5 w-5 flex-shrink-0","aria-hidden":"true"})],8,Fe)])])])]))),128))]),e("nav",ze,[e("div",Be,[e("button",{disabled:d.value.isOnFirst(),class:"btn btn--default",onClick:s[1]||(s[1]=x(o=>T("prev"),["prevent"]))},n(t(a)("pagination.prev")),9,De),e("button",{disabled:d.value.isOnLast(),class:"ml-3 btn btn--default",onClick:s[2]||(s[2]=x(o=>T("next"),["prevent"]))},n(t(a)("pagination.next")),9,Ne)])])],2)):y("",!0),d.value?(l(),i("div",Re,[e("div",{class:U(["mx-auto max-w-6xl px-4 md:px-6 lg:px-8 transition-all duration-200",{"pointer-events-none opacity-75":_.value}])},[e("div",Se,[e("div",Oe,[e("table",Ge,[e("thead",null,[e("tr",null,[e("th",Ae,n(t(a)("users.username")),1),e("th",Ie,n(t(a)("pagination.actions")),1)])]),e("tbody",Je,[(l(!0),i(v,null,X(j.value,o=>(l(),i("tr",{key:o.username,class:"bg-context-50 dark:bg-context-800"},[e("td",Ke,[e("div",Qe,n(o.username),1)]),e("td",We,[e("a",{class:"cursor-pointer transition-all duration-500 text-context-500 dark:text-context-400 hover:text-context-700 hover:dark:text-context-300",onClick:x(k=>F(o),["prevent"])},[m(t(H),{class:"w-5 h-5","aria-hidden":"true"})],8,Xe),e("a",{class:"cursor-pointer transition-all duration-500 text-context-500 dark:text-context-400 hover:text-context-700 hover:dark:text-context-300",onClick:x(k=>z(o),["prevent"])},[m(t(Y),{class:"w-5 h-5","aria-hidden":"true"})],8,Ye)])]))),128))])]),e("nav",Ze,[e("div",He,[e("p",{class:"text-sm text-context-700 dark:text-context-200",innerHTML:t(a)("pagination.text",{from:d.value.getFrom(),to:d.value.getTo(),of:d.value.getTotal()})},null,8,et)]),e("div",tt,[e("button",{disabled:d.value.isOnFirst(),class:"btn btn--default",onClick:s[3]||(s[3]=x(o=>T("prev"),["prevent"]))},n(t(a)("pagination.prev")),9,st),e("button",{disabled:d.value.isOnLast(),class:"ml-3 btn btn--default",onClick:s[4]||(s[4]=x(o=>T("next"),["prevent"]))},n(t(a)("pagination.next")),9,at)])])])])],2)])):y("",!0)],64))]),m(t(_e),{as:"template",show:u.value!==null},{default:g(()=>[m(t(xe),{as:"div",class:U(["relative z-10",{"pointer-events-none":_.value}]),onClose:E},{default:g(()=>[m(t(Z),{as:"template",enter:"ease-out duration-300","enter-from":"opacity-0","enter-to":"opacity-100",leave:"ease-in duration-200","leave-from":"opacity-100","leave-to":"opacity-0"},{default:g(()=>[ot]),_:1}),e("div",nt,[e("div",lt,[m(t(Z),{as:"template",enter:"ease-out duration-300","enter-from":"opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95","enter-to":"opacity-100 translate-y-0 sm:scale-100",leave:"ease-in duration-200","leave-from":"opacity-100 translate-y-0 sm:scale-100","leave-to":"opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95"},{default:g(()=>[m(t(me),{class:"bg-context-50 dark:bg-context-900 relative transform overflow-hidden rounded-lg px-4 pt-5 pb-4 text-left shadow-xl transition-all sm:my-8 sm:w-full sm:max-w-lg sm:p-6"},{default:g(()=>[e("form",{onSubmit:x(te,["prevent"])},[e("div",null,[e("div",it,[m(t(ve),{as:"h3",class:"text-lg font-medium leading-6 text-context-900 dark:text-context-100"},{default:g(()=>[u.value===""?(l(),i(v,{key:0},[h(n(t(a)("users.modal.createTitle")),1)],64)):(l(),i(v,{key:1},[h(n(t(a)("users.modal.editTitle",{user:u.value})),1)],64))]),_:1}),e("div",dt,[e("div",ct,[u.value===""?(l(),i("div",ut,[e("label",pt,n(t(a)("users.username")),1),e("div",xt,[$(e("input",{id:"new_username","onUpdate:modelValue":s[5]||(s[5]=o=>p.value.new_username=o),name:"new_username",type:"text",autocomplete:"off",required:u.value==="",class:"form-input",placeholder:t(a)("users.username")},null,8,mt),[[M,p.value.new_username]])])])):y("",!0),e("div",null,[e("label",vt,n(t(a)("users.hubPassword")),1),e("div",_t,[$(e("input",{id:"hub_password","onUpdate:modelValue":s[6]||(s[6]=o=>p.value.hub_password=o),name:"hub_password",type:"password",autocomplete:"hub_password",required:u.value==="",class:"form-input",placeholder:t(a)("users.hubPassword")},null,8,ht),[[M,p.value.hub_password]])]),e("div",ft,[u.value===""?(l(),i(v,{key:0},[],64)):(l(),i(v,{key:1},[h(n(t(a)("users.emptyPasswordHint")),1)],64))])]),e("div",null,[e("label",bt,n(t(a)("users.smtpPassword")),1),e("div",wt,[$(e("input",{id:"smtp_password","onUpdate:modelValue":s[7]||(s[7]=o=>p.value.smtp_password=o),name:"smtp_password",type:"password",autocomplete:"smtp_password",class:"form-input",placeholder:t(a)("users.smtpPassword")},null,8,kt),[[M,p.value.smtp_password]])]),e("div",yt,[u.value===""?(l(),i(v,{key:0},[h(n(t(a)("users.emptySmtpPasswordHint")),1)],64)):(l(),i(v,{key:1},[h(n(t(a)("users.emptyPasswordHint")),1)],64))])])])])])]),e("div",gt,[e("button",{type:"button",class:"btn btn--default",onClick:x(E,["prevent"])},n(t(a)("users.modal.cancel")),9,Ct),e("button",Tt,n(t(a)("users.modal.submit")),1)])],40,rt)]),_:1})]),_:1})])])]),_:1},8,["class"])]),_:1},8,["show"]),V.value?(l(),ce(ue,{key:0,to:"#header-search"},[e("form",{class:"flex w-full md:ml-0",method:"GET",onSubmit:s[9]||(s[9]=x(o=>w(1),["prevent"]))},[e("label",$t,n(t(a)("inbox.search")),1),e("div",jt,[e("div",Pt,[m(t(pe),{class:"h-5 w-5","aria-hidden":"true"})]),$(e("input",{id:"search-field","onUpdate:modelValue":s[8]||(s[8]=o=>C.value.search=o),name:"search-field",class:"block h-full w-full py-2 pl-8 pr-3 sm:text-sm focus:outline-none focus:ring-0 border-transparent focus:border-transparent placeholder-context-500 bg-context-50 dark:bg-context-900 text-context-900 dark:text-context-100",placeholder:t(a)("users.search"),type:"search"},null,8,Ut),[[M,C.value.search]])])],32)])):y("",!0)],64)}}});export{Ft as default};
