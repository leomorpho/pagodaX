(()=>{function g(i){let r=[{question:"What is the capital of France?",options:["Paris","London","Berlin","Madrid"],answer:"Paris"},{question:"What is 2 + 2?",options:["3","4","5","6"],answer:"4"}],t=i.querySelector("#js-quiz-container");if(!t)return;t.setAttribute("data-initialized","true"),t.innerHTML="",t.className="max-w-xl mx-auto bg-gray-100 dark:bg-gray-700 text-black dark:text-white shadow-md rounded-lg p-8 mt-5",r.forEach((a,s)=>{let n=document.createElement("div");n.className="mb-4 last:mb-0";let e=document.createElement("p");e.textContent=`${s+1}. ${a.question}`,e.className="font-semibold text-lg mb-2",n.appendChild(e);let c=document.createElement("div");c.className="pl-4",a.options.forEach(h=>{let m=document.createElement("div");m.className="flex items-center mb-1";let d=document.createElement("input");d.type="radio",d.name=`question${s}`,d.value=h,d.className="option-input mr-2";let u=document.createElement("label");u.appendChild(document.createTextNode(h)),u.className="select-none",m.appendChild(d),m.appendChild(u),c.appendChild(m)}),n.appendChild(c),t.appendChild(n)});let o=document.createElement("button");o.textContent="Submit Answers",o.className="mt-4 px-4 py-2 bg-blue-500 hover:bg-blue-700 text-white font-bold rounded cursor-pointer",o.addEventListener("click",f.bind(null,r,i)),t.appendChild(o)}function f(i,r){let t=0;i.forEach((o,a)=>{let s=r.querySelector(`input[name="question${a}"]:checked`);s&&s.value===o.answer&&t++}),alert(`You scored ${t}/${i.length}`)}window.initializeJS=function(r){console.log("initialize js"),b(),window.darkModeSwitchersInitialized=window.darkModeSwitchersInitialized||!1,v();let t=r||document,o=t.querySelector("#js-quiz-container");o&&!o.hasAttribute("data-initialized")&&g(t),"serviceWorker"in navigator&&navigator.serviceWorker.register("/service-worker.js",{type:"module"}).then(function(a){console.log("Service Worker registered with scope:",a.scope)}).catch(function(a){console.log("Service Worker registration failed:",a)})};function b(){let i=document.querySelector('meta[name="viewport"]');function r(){window.innerWidth<1024?i.setAttribute("content","width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=no"):i.setAttribute("content","width=device-width, initial-scale=1.0")}r(),window.addEventListener("resize",r)}var p="lightmode",l="darkmode",w="dark";function v(){if(window.darkModeSwitchersInitialized)return;let i=document.querySelectorAll(".theme-toggle"),r=document.querySelectorAll(".theme-toggle-dark-icon"),t=document.querySelectorAll(".theme-toggle-light-icon");function o(){let n=document.documentElement.getAttribute("data-theme")===l;r.forEach(e=>{e.classList.toggle("hidden",n)}),t.forEach(e=>{e.classList.toggle("hidden",!n)})}function a(){let n=localStorage.getItem("color-theme"),e=window.matchMedia("(prefers-color-scheme: dark)").matches,c=p;n?c=n:e&&(c=l),document.documentElement.setAttribute("data-theme",c),document.documentElement.classList.toggle(w,c===l),o()}a();function s(){let n=document.documentElement.getAttribute("data-theme")===l,e=document.documentElement.style;n?e.setProperty("--brightness-hover","var(--brightness-hover-dark)"):e.setProperty("--brightness-hover","var(--brightness-hover-light)")}i.forEach(n=>{n.addEventListener("click",()=>{let e=document.documentElement.getAttribute("data-theme")===l?p:l;document.documentElement.setAttribute("data-theme",e),document.documentElement.classList.toggle(w,e===l),localStorage.setItem("color-theme",e),o(),s()})}),window.darkModeSwitchersInitialized=!0}})();
//# sourceMappingURL=vanilla_bundle.js.map
