package telegram

const msgHelp = `I can save and keep your pages. Also I can offer you them to read.

In order to save the page, just send me link.

To get random page - send me command /rnd.

Caution!`

const msgHello = "Hi there! \n\n" + msgHelp

const (
	msgUnknownCommand = "Я не знаю эту команду"
	msgNoSavedPages   = "У вас нет сохраненных страниц"
	msgSaved          = "Сохранено!"
	msgAlreadyExists  = "Эта страница уже есть в вашем списке!"
)
