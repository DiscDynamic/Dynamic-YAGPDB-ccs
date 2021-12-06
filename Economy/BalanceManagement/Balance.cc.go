{{/*
        Made by Ranger (765316548516380732)

    Trigger Type: `Command`
    Trigger: `Balance`
©️ Ranger 2021
MIT License
*/}}

{{/* Only edit below if you know what you're doing (: rawr */}}

{{$a := ""}}
{{$cash := ""}}
{{$bank := ""}}
{{$EconomySymbol := ""}}
{{$b := .User.ID}}
{{if not (dbGet $b "EconomyInfo")}}
    {{dbSet .User.ID "EconomyInfo" (sdict "cash" 0 "bank" 0)}}
{{end}}
{{with (dbGet 0 "EconomySettings")}}
	{{$a = sdict .Value}}
	{{$EconomySymbol = $a.EconomySymbol}}
	{{with (dbGet $b "EconomyInfo")}}
        {{$a = sdict .Value}}
		{{$cash = $a.cash}}
		{{$bank = $a.bank}}
		{{$balanceEmbed := (cembed
            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
            "description" (print "Your balance")
            "fields" (cslice 
                (sdict "name" "Cash" "value" (print $EconomySymbol (toInt $cash)) "inline" true)
                (sdict "name" "Bank" "value" (print $EconomySymbol (toInt $bank)) "inline" true)
                (sdict "name" "Networth" "value" (print $EconomySymbol (toString (add (toInt $cash) (toInt $bank)))) "inline" true))
            "color" 0x00ff7b
            "timestamp" currentTime
            )}}
		{{sendMessage nil $balanceEmbed}}
	{{end}}
{{end}}