function copyLink() {
 	const copy_icon = document.getElementById('copy_icon');

 	const short_link = document.getElementById('short_link');
 	const link = short_link.innerHTML;

	navigator.clipboard.writeText(link).then(function() {
        console.log('Link copied: ' + link);
 		copy_icon.style.fill = "green";
 		setTimeout(() => {
			copy_icon.style.fill = "black";
		}, 2000);
    }).catch(function(error) {
        console.error('Error copying text: ', error);
		copy_icon.style.fill = "red";
		setTimeout(() => {
			copy_icon.style.fill = "black";
		}, 2000);
    });
} 
