// сheck if navigator.clipboard is working in current browser

// if (navigator.clipboard) {
// 	console.log('navigator.clipboard is working')
//
// 	const header = document.createElement("h1");        // создаем заголовок <h1>
// 	const  headerText = document.createTextNode("Hello World"); // создаем текстовый узел
// 	header.appendChild( headerText); // добавляем в элемент h1 текстовый узел
//
// 	const main = document.getElementById('main');
// 	document.body.insertBefore(header, main);
// } else {
// 	console.log('navigator.clipboard NOT working')
//
// 	const header = document.createElement("h1");        // создаем заголовок <h1>
// 	const  headerText = document.createTextNode("Not Dave"); // создаем текстовый узел
// 	header.appendChild( headerText); // добавляем в элемент h1 текстовый узел
//
// 	const main = document.getElementById('main');
// 	document.body.insertBefore(header, main);
// }


// if navigator.clipboard is working in current browser you may use this method to copy link to buffer

// function copyLink() {
//  	const copy_icon = document.getElementById('copy_icon');
//
//  	const short_link = document.getElementById('short_link');
//  	const link = short_link.innerHTML;
//
// 	navigator.clipboard.writeText(link).then(function() {
//         console.log('Link copied: ' + link);
//  		copy_icon.style.fill = "green";
//  		setTimeout(() => {
// 			copy_icon.style.fill = "black";
// 		}, 2000);
//     }).catch(function(error) {
//         console.error('Error copying text: ', error);
// 		copy_icon.style.fill = "red";
// 		setTimeout(() => {
// 			copy_icon.style.fill = "black";
// 		}, 2000);
//     });
// }


// the best miltiplatform method to copy link to buffer 

const copy_icon = document.getElementById('copy_icon');
const short_link = document.getElementById('short_link')

copy_icon.addEventListener('click', (evt) => {
    const input = document.createElement('input');
    const link = short_link.innerHTML;
    const yPosition = window.pageYOffset || document.documentElement.scrollTop;

    // Скрываем поле за краями экрана в зависимости от направления текста в текущей локали
    input.style.position = 'absolute';
    input.style['right'] = '-9999px';

    // Предотвращаем срабатывание зума на iOS
    input.style.fontSize = '16px';

    // Предотвращаем скролл к элементу
    input.style.top = `${yPosition}px`;

    // Не допускаем появления клавиатуры на мобильных девайсах
    input.setAttribute('readonly', '');

    // Вставляем элемент в DOM
    document.body.appendChild(input);
    input.value = link;

    // Помещаем поле в фокус
    input.focus();
    // Выделяем текст в поле
    input.select();
    // Копируем текст в поле обмена
    document.execCommand('copy');

    // Удаляем фейковый элемент
    document.body.removeChild(input);
});
