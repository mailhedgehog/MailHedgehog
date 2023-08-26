import{b as X,i as Z,J as tt,u as et,l as g,e as st,E as ot,m as nt,p as at,o as c,c as i,a as t,k as s,F as x,j as _,t as o,n as D,q as z,s as u,v as f,x as m,f as J,g as ct,K as it,G as lt}from"./index.99471e1d.js";import{h as $,r as rt,P as j}from"./pagination.4f0e65cd.js";import{r as M}from"./TrashIcon.a64c37b7.js";import{r as K}from"./EyeIcon.d2dee620.js";const dt={class:"lg:-mt-px bg-context-50 dark:bg-context-900 shadow dark:shadow-context-500"},xt={class:"px-4 sm:px-6 lg:mx-auto lg:max-w-6xl lg:px-8"},ut={class:"py-6 flex items-center justify-between"},pt={class:"ml-3 text-xl md:text-2xl font-bold leading-7 text-context-900 dark:text-context-100 truncate sm:leading-9"},mt={class:"hidden md:inline"},ht={class:"mt-8"},_t={key:0,class:""},vt={class:"ml-3 text-lg md:text-xl text-center text-context-900 dark:text-context-100 select-none"},ft={role:"list",class:"mt-2 divide-y divide-context-200 overflow-hidden shadow dark:shadow-context-500"},bt={class:"block bg-context-50 dark:bg-context-800 px-4 py-4 hover:bg-context-50"},kt={class:"flex items-center space-x-4"},gt={class:"flex flex-1 space-x-2 truncate"},yt={class:"flex flex-col truncate text-sm text-context-500 dark:text-context-400"},wt={key:0,class:"truncate"},Yt={class:"font-medium text-context-900 dark:text-context-100 mr-2"},$t={key:1,class:"truncate"},Ct={class:"font-medium text-context-900 dark:text-context-100 mr-2"},Dt={class:"font-medium text-context-900 dark:text-context-100 mr-2"},jt={class:"font-medium text-context-900 dark:text-context-100 mr-2"},Mt=["datetime"],Tt={class:"space-y-2"},Et=["onClick"],qt=["onClick"],Vt={class:"flex items-center justify-between border-t border-context-200 bg-context-50 dark:bg-context-800 px-4 py-3","aria-label":"Pagination"},Ft={class:"flex flex-1 justify-between"},Nt=["disabled"],Pt=["disabled"],Lt={key:1,class:"hidden md:block"},Bt={class:"mt-2 flex flex-col"},Ot={class:"min-w-full overflow-hidden overflow-x-auto align-middle shadow md:rounded-lg"},Rt={class:"min-w-full divide-y divide-context-200"},It={class:"bg-context-50 dark:bg-context-700 px-6 py-3 text-left text-sm font-semibold text-context-900 dark:text-context-100",scope:"col"},St={class:"bg-context-50 dark:bg-context-700 px-6 py-3 text-left text-sm font-semibold text-context-900 dark:text-context-100",scope:"col"},At={class:"bg-context-50 dark:bg-context-700 px-6 py-3 text-left text-sm font-semibold text-context-900 dark:text-context-100",scope:"col"},Gt={class:"whitespace-nowrap bg-context-50 dark:bg-context-700 px-6 py-3 text-left text-sm font-semibold text-context-900 dark:text-context-100",scope:"col"},Ht={class:"bg-context-50 dark:bg-context-700 px-6 py-3 text-sm font-semibold text-context-900 dark:text-context-100 text-right",scope:"col"},Ut={class:"divide-y divide-context-200 bg-context-50 dark:bg-context-800"},zt={class:"max-w-[12rem] whitespace-nowrap px-6 py-4 text-sm text-context-900 dark:text-context-100"},Jt={class:"truncate"},Kt={class:"truncate"},Qt=["href"],Wt={key:1},Xt={class:"max-w-[12rem] whitespace-nowrap px-6 py-4 text-sm text-context-900 dark:text-context-100"},Zt={class:"truncate"},te={class:"truncate"},ee=["href"],se={key:1},oe={class:"w-full max-w-0 whitespace-nowrap truncate px-6 py-4 text-sm text-context-900 dark:text-context-100"},ne={class:"whitespace-nowrap px-6 py-4 text-sm text-context-500 dark:text-context-400"},ae=["datetime"],ce={class:"whitespace-nowrap px-6 py-4 text-sm text-right flex justify-end space-x-1"},ie=["onClick"],le=["onClick"],re={class:"flex items-center justify-between border-t border-context-200 bg-context-50 dark:bg-context-800 px-4 py-3 sm:px-6","aria-label":"Pagination"},de={class:"hidden sm:block"},xe=["innerHTML"],ue={class:"flex flex-1 justify-between sm:justify-end"},pe=["disabled"],me=["disabled"],he={for:"search-field",class:"sr-only"},_e={class:"relative w-full text-context-400 focus-within:text-context-600"},ve={class:"pointer-events-none absolute inset-y-0 left-0 flex items-center","aria-hidden":"true"},fe=["placeholder"],Ye={__name:"Inbox",setup(be){const{t:n}=X(),l=Z("MailHedgehog"),Q=tt(),C=et(),b=g({page:1,per_page:25,search:""}),T=g(!1),d=g(!1),y=g([]),r=g(new j),E=st(()=>C.getters.getUser),h=(p=null)=>{d.value=!0,p&&(b.value.page=p),l==null||l.request().get("emails",{params:b.value}).then(a=>{var Y,e,v,F,N,P,L,B,O,R,I,S,A,G,H,U;(Y=a.data)!=null&&Y.data?y.value=(e=a.data)==null?void 0:e.data:y.value=[],(F=(v=a.data)==null?void 0:v.meta)!=null&&F.pagination?r.value=new j((P=(N=a.data)==null?void 0:N.meta)==null?void 0:P.pagination.current_page,(B=(L=a.data)==null?void 0:L.meta)==null?void 0:B.pagination.per_page,(R=(O=a.data)==null?void 0:O.meta)==null?void 0:R.pagination.last_page,(S=(I=a.data)==null?void 0:I.meta)==null?void 0:S.pagination.from,(G=(A=a.data)==null?void 0:A.meta)==null?void 0:G.pagination.to,(U=(H=a.data)==null?void 0:H.meta)==null?void 0:U.pagination.total):r.value=new j}).catch(a=>{l.onResponseError(a,"Response Error")}).finally(()=>{d.value=!1})};let w=null;ot(()=>b.value.search,()=>{w&&(clearTimeout(w),w=null),w=setTimeout(()=>{h(1)},500)},{deep:!0}),nt(()=>{h(1),T.value=!0}),l.$on("new_message",()=>h());const k=(p="next")=>{h(r.value.getPageFromDirection(p))},W=()=>{C.dispatch("confirmDialog/confirm").then(()=>{d.value=!0,l==null||l.request().delete("emails").then(()=>{h(1),l==null||l.success(n("inbox.cleared"))}).catch(()=>{d.value=!1})})},q=p=>{Q.push({name:"email",params:{id:p}})},V=p=>{C.dispatch("confirmDialog/confirm").then(()=>{d.value=!0,l==null||l.request().delete(`emails/${p}`).then(()=>{r.value.count()===0?k("prev"):h(),l==null||l.success(n("email.deleted"))}).catch(()=>{d.value=!1})})};return(p,a)=>{const Y=at("tooltip");return c(),i(x,null,[t("div",dt,[t("div",xt,[t("div",ut,[t("h1",pt,[s(E)?(c(),i(x,{key:0},[_(o(s(n)("inbox.hello",{msg:s(E).username})),1)],64)):(c(),i(x,{key:1},[_(o(s(n)("inbox.pageTitle")),1)],64))]),t("div",{class:D(["flex justify-end ml-4 transition-all duration-200",{"pointer-events-none opacity-75":d.value}])},[r.value&&!r.value.isEmpty()?z((c(),i("button",{key:0,type:"button",class:"btn btn--default whitespace-nowrap",onClick:a[0]||(a[0]=u(e=>W(),["prevent"]))},[f(s(M),{class:"w-4 h-4 md:mr-2"}),t("span",mt,o(s(n)("inbox.clear")),1)])),[[Y,s(n)("inbox.clear")]]):m("",!0)],2)])])]),t("div",ht,[r.value.isEmpty()?(c(),i("div",_t,[t("h3",vt,[d.value?(c(),i(x,{key:0},[_(o(s(n)("pagination.requesting")),1)],64)):(c(),i(x,{key:1},[_(o(s(n)("inbox.empty")),1)],64))])])):(c(),i(x,{key:1},[r.value?(c(),i("div",{key:0,class:D(["shadow dark:shadow-context-500 md:hidden",{"pointer-events-none opacity-75":d.value}])},[t("ul",ft,[(c(!0),i(x,null,J(y.value,e=>(c(),i("li",{key:e.id},[t("div",bt,[t("span",kt,[t("span",gt,[t("span",yt,[e.from?(c(),i("span",wt,[t("span",Yt,o(s(n)("email.from")),1),_(" "+o(e.from.name)+"("+o(e.from.email)+") ",1)])):m("",!0),e.to[0]?(c(),i("span",$t,[t("span",Ct,o(s(n)("email.to")),1),_(" "+o(e.to[0].name)+"("+o(e.to[0].email)+") ",1)])):m("",!0),t("span",null,[t("span",Dt,o(s(n)("email.subject")),1),_(" "+o(e.subject),1)]),t("span",null,[t("span",jt,o(s(n)("email.received_at")),1),s($)(e.received_at,"YYYY-MM-DD HH:mm:ss").isValid()?(c(),i("time",{key:0,datetime:e.received_at},o(s($)(e.received_at,"YYYY-MM-DD HH:mm:ss").fromNow()),9,Mt)):m("",!0)])])]),t("div",Tt,[t("a",{class:"cursor-pointer block transition-all duration-500 text-context-500 dark:text-context-400 hover:text-context-700 hover:dark:text-context-300",onClick:u(v=>q(e.id),["prevent"])},[f(s(K),{class:"h-5 w-5 flex-shrink-0","aria-hidden":"true"})],8,Et),t("a",{class:"cursor-pointer block transition-all duration-500 text-context-500 dark:text-context-400 hover:text-context-700 hover:dark:text-context-300",onClick:u(v=>V(e.id),["prevent"])},[f(s(M),{class:"h-5 w-5 flex-shrink-0","aria-hidden":"true"})],8,qt)])])])]))),128))]),t("nav",Vt,[t("div",Ft,[t("button",{disabled:r.value.isOnFirst(),class:"btn btn--default",onClick:a[1]||(a[1]=u(e=>k("prev"),["prevent"]))},o(s(n)("pagination.prev")),9,Nt),t("button",{disabled:r.value.isOnLast(),class:"ml-3 btn btn--default",onClick:a[2]||(a[2]=u(e=>k("next"),["prevent"]))},o(s(n)("pagination.next")),9,Pt)])])],2)):m("",!0),r.value?(c(),i("div",Lt,[t("div",{class:D(["mx-auto max-w-6xl px-4 md:px-6 lg:px-8 transition-all duration-200",{"pointer-events-none opacity-75":d.value}])},[t("div",Bt,[t("div",Ot,[t("table",Rt,[t("thead",null,[t("tr",null,[t("th",It,o(s(n)("email.from")),1),t("th",St,o(s(n)("email.to")),1),t("th",At,o(s(n)("email.subject")),1),t("th",Gt,o(s(n)("email.received_at")),1),t("th",Ht,o(s(n)("pagination.actions")),1)])]),t("tbody",Ut,[(c(!0),i(x,null,J(y.value,e=>(c(),i("tr",{key:e.id,class:"bg-context-50 dark:bg-context-800"},[t("td",zt,[e.from?(c(),i(x,{key:0},[t("div",Jt,o(e.from.name),1),t("div",Kt,[t("a",{href:`mailto:${e.from.email}`,class:"text-context-500 dark:text-context-400 hover:text-context-700 hover:dark:text-context-200 transition-all duration-500"},o(e.from.email),9,Qt)])],64)):(c(),i("div",Wt,o(s(n)("email.notAvailable")),1))]),t("td",Xt,[e.to[0]?(c(),i(x,{key:0},[t("div",Zt,o(e.to[0].name),1),t("div",te,[t("a",{href:`mailto:${e.to[0].email}`,class:"text-context-500 dark:text-context-400 hover:text-context-700 hover:dark:text-context-200 transition-all duration-500"},o(e.to[0].email),9,ee)])],64)):(c(),i("div",se,o(s(n)("email.notAvailable")),1))]),t("td",oe,o(e.subject),1),t("td",ne,[s($)(e.received_at,"YYYY-MM-DD HH:mm:ss").isValid()?(c(),i("time",{key:0,datetime:e.received_at},o(s($)(e.received_at,"YYYY-MM-DD HH:mm:ss").fromNow()),9,ae)):m("",!0)]),t("td",ce,[t("a",{class:"cursor-pointer transition-all duration-500 text-context-500 dark:text-context-400 hover:text-context-700 hover:dark:text-context-300",onClick:u(v=>q(e.id),["prevent"])},[f(s(K),{class:"w-5 h-5","aria-hidden":"true"})],8,ie),t("a",{class:"cursor-pointer transition-all duration-500 text-context-500 dark:text-context-400 hover:text-context-700 hover:dark:text-context-300",onClick:u(v=>V(e.id),["prevent"])},[f(s(M),{class:"w-5 h-5","aria-hidden":"true"})],8,le)])]))),128))])]),t("nav",re,[t("div",de,[t("p",{class:"text-sm text-context-700 dark:text-context-200",innerHTML:s(n)("pagination.text",{from:r.value.getFrom(),to:r.value.getTo(),of:r.value.getTotal()})},null,8,xe)]),t("div",ue,[t("button",{disabled:r.value.isOnFirst(),class:"btn btn--default",onClick:a[3]||(a[3]=u(e=>k("prev"),["prevent"]))},o(s(n)("pagination.prev")),9,pe),t("button",{disabled:r.value.isOnLast(),class:"ml-3 btn btn--default",onClick:a[4]||(a[4]=u(e=>k("next"),["prevent"]))},o(s(n)("pagination.next")),9,me)])])])])],2)])):m("",!0)],64))]),T.value?(c(),ct(lt,{key:0,to:"#header-search"},[t("form",{class:"flex w-full md:ml-0",method:"GET",onSubmit:a[6]||(a[6]=u(e=>h(1),["prevent"]))},[t("label",he,o(s(n)("inbox.search")),1),t("div",_e,[t("div",ve,[f(s(rt),{class:"h-5 w-5","aria-hidden":"true"})]),z(t("input",{id:"search-field","onUpdate:modelValue":a[5]||(a[5]=e=>b.value.search=e),name:"search-field",class:"block h-full w-full py-2 pl-8 pr-3 sm:text-sm focus:outline-none focus:ring-0 border-transparent focus:border-transparent placeholder-context-500 bg-context-50 dark:bg-context-900 text-context-900 dark:text-context-100",placeholder:s(n)("inbox.search"),type:"search"},null,8,fe),[[it,b.value.search]])])],32)])):m("",!0)],64)}}};export{Ye as default};