{{/*
		Made by ranger_4297 (765316548516380732)

	Trigger Type: `Regex`
	Trigger: `\A(-|<@!?204255221017214977>\s*)(coin-?flip|cf|flip-?coin|fc)(\s+|\z)`

	©️ RhykerWells 2020-Present
	GNU, GPLV3 License
	Repository: https://github.com/Ranger-4297/YAGPDB-ccs
*/}}


{{/* Only edit below if you know what you're doing (: rawr */}}

{{/* Initiates variables */}}
{{$userID := .User.ID}}
{{$successColor := 0x00ff7b}}
{{$errorColor := 0xFF0000}}
{{$prefix := .ServerPrefix}}

{{/* Flip */}}

{{/* Response */}}
{{$embed := sdict "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "1024")) "timestamp" currentTime "color" $errorColor}}
{{$economySettings := (dbGet 0 "EconomySettings").Value}}
{{if not $economySettings}}
	{{$embed.Set "description" (print "No `Settings` database found.\nPlease set it up with the default values using `" .ServerPrefix "server-set default`")}}
	{{sendMessage nil (cembed $embed)}}
	{{return}}
{{end}}
{{$symbol := $economySettings.symbol}}
{{$betMax := $economySettings.betMax | toInt}}
{{$incomeCooldown := $economySettings.incomeCooldown | toInt}}
{{if not .CmdArgs}}
    {{$embed.Set "description" (print "No `bet` argument provided.\nSyntax is `" .Cmd " <Bet:Amount>`")}}
	{{sendMessage nil (cembed $embed)}}
    {{return}}
{{end}}
{{$bal := toInt (dbGet $userID "cash").Value}}
{{$bet := index .CmdArgs 0 | str | lower}}
{{if not (or (toInt $bet) (eq $bet "all" "max"))}}
    {{$embed.Set "description" (print "Invalid `Bet` argument provided.\nSyntax is `" .Cmd " <Bet:Amount>`")}}
    {{sendMessage nil (cembed $embed)}}
    {{return}}
{{end}}
{{if eq $bet "all"}}
	{{$bet = $bal}}
{{else if eq $bet "max"}}
	{{$bet = $betMax}}
{{end}}
{{if le ($bet = toInt $bet) 0}}
    {{$embed.Set "description" (print "Invalid `Bet` argument provided.\nSyntax is `" .Cmd " <Bet:Amount>`")}}
    {{sendMessage nil (cembed $embed)}}
    {{return}}
{{end}}
{{if gt $bet $bal}}
    {{$embed.Set "description" (print "You can't bet more than you have!")}}
    {{sendMessage nil (cembed $embed)}}
    {{return}}
{{end}}
{{if gt $bet $betMax}}
    {{$embed.Set "description" (print "You can't bet more than " $symbol $betMax)}}
    {{sendMessage nil (cembed $embed)}}
    {{return}}
{{end}}
{{$side := index .CmdArgs 1 | str | lower}}
{{if not $side}}
	{{$embed.Set "description" (print "No `Side` argument provided.\nSyntax is `" $.Cmd " <Side:Head/Tails> <Bet:Amount>`")}}
	{{sendMessage nil (cembed $embed)}}
    {{return}}
{{end}}
{{if not (`\A(t(ails?)?|h(eads?)?)(\s+|\z)` $side)}}
	{{$embed.Set "description" (print "Invalid `Side` argument provided.\nSyntax is `" $.Cmd " <Bet:Amount> <Side:Heads/Tails>`")}}
	{{sendMessage nil (cembed $embed)}}
    {{return}}
{{end}}
{{if eq (slice $side 0 1) "t"}}
	{{$side = "tails"}}
{{else}}
	{{$side = "heads"}}
{{end}}
{{if ($cooldown := dbGet $userID "coinflipCooldown")}}
	{{$embed.Set "description" (print "This command is on cooldown for " (humanizeDurationSeconds ($cooldown.ExpiresAt.Sub currentTime)))}}
	{{sendMessage nil (cembed $embed)}}
    {{return}}
{{end}}
{{dbSetExpire $userID "coinflipCooldown" "cooldown" $incomeCooldown}}
{{if and ($int := randInt 1 3) eq $int 1}}
	{{$bal = add $bal $bet}}
	{{$embed.Set "description" (print "You flipped " $side " and won " $symbol (humanizeThousands $bet))}}
	{{$embed.Set "color" $successColor}}
{{else}}
	{{$bal = sub $bal $bet}}
	{{$embed.Set "description" (print "You flipped " $side " and lost.")}}
	{{$embed.Set "color" $errorColor}}
{{end}}
{{dbSet $userID "cash" $bal}}
{{sendMessage nil (cembed $embed)}}