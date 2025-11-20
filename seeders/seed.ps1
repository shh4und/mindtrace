# PowerShell script to mimic seed.sh behavior on Windows
$ErrorActionPreference = 'Stop'
Set-StrictMode -Version Latest

$ApiBase = 'http://localhost:8181/api/v1'

function New-Professional {
    param (
        [string]$Nome,
        [string]$Email,
        [string]$Especialidade,
        [string]$Registro,
        [string]$Cpf,
        [string]$Contato,
        [string]$DataNascimento
    )

    Write-Host "Registering professional: $Nome"

    $payload = @{
        nome                = $Nome
        email               = $Email
        senha               = 'Password123!'
        data_nascimento     = $DataNascimento
        especialidade       = $Especialidade
        registro_profissional = $Registro
        cpf                 = $Cpf
        contato             = $Contato
    } | ConvertTo-Json -Compress

    try {
        $response = Invoke-RestMethod -Method Post -Uri "$ApiBase/profissionais/registrar" -ContentType 'application/json' -Body $payload -ErrorAction Stop
        if ($response.mensagem) {
            Write-Host $response.mensagem
        } elseif ($response.erro) {
            Write-Host $response.erro
        } else {
            Write-Host 'Request sent'
        }
    } catch {
        Write-Warning "Failed to register professional ${Nome}: $_"
    }
}

function New-Patient {
    param (
        [string]$Nome,
        [string]$Email,
        [string]$Cpf,
        [bool]$Dependente,
        [string]$DataNascimento,
        [string]$Contato,
        [string]$NomeResponsavel,
        [string]$ContatoResponsavel
    )

    Write-Host "Registering patient: $Nome"

    $payload = @{
        nome               = $Nome
        email              = $Email
        senha              = 'Password123!'
        cpf                = $Cpf
        dependente         = $Dependente
        data_nascimento    = $DataNascimento
        contato            = $Contato
    }

    if ($Dependente -and $NomeResponsavel) {
        $payload.nome_responsavel = $NomeResponsavel
        $payload.contato_responsavel = $ContatoResponsavel
    }

    $json = $payload | ConvertTo-Json -Compress

    try {
        $response = Invoke-RestMethod -Method Post -Uri "$ApiBase/pacientes/registrar" -ContentType 'application/json' -Body $json -ErrorAction Stop
        if ($response.mensagem) {
            Write-Host $response.mensagem
        } elseif ($response.erro) {
            Write-Host $response.erro
        } else {
            Write-Host 'Request sent'
        }
    } catch {
        Write-Warning "Failed to register patient ${Nome}: $_"
    }
}

Write-Host "API Base: $ApiBase" -ForegroundColor Cyan

Write-Host 'Registering professionals...'
New-Professional 'Dr. João Silva' 'joao.silva@mindtrace.com' 'Psicologia Clínica' 'CRP-01/12345' '12345678901' '11999990001' '1985-04-15T00:00:00Z'
New-Professional 'Dra. Maria Santos' 'maria.santos@mindtrace.com' 'Psiquiatria' 'CRM-02/67890' '12345678902' '11999990002' '1980-08-22T00:00:00Z'
New-Professional 'Dr. Pedro Oliveira' 'pedro.oliveira@mindtrace.com' 'Terapia Cognitivo-Comportamental' 'CRP-03/54321' '12345678903' '11999990003' '1988-12-10T00:00:00Z'
New-Professional 'Dra. Ana Paula Costa' 'ana.costa@mindtrace.com' 'Psicologia Infantil' 'CRP-04/98765' '12345678920' '11999990004' '1990-03-28T00:00:00Z'
New-Professional 'Dr. Carlos Eduardo Lima' 'carlos.lima@mindtrace.com' 'Neuropsicologia' 'CRP-05/11111' '12345678921' '11999990005' '1982-07-05T00:00:00Z'

Write-Host 'Registering patients...'
New-Patient 'Ana Costa' 'ana.paciente@mindtrace.com' '12345678904' $false '1995-03-15T00:00:00Z' '11777770001' '' ''
New-Patient 'Bruno Lima' 'bruno.paciente@mindtrace.com' '12345678905' $true '2010-07-22T00:00:00Z' '11777770002' 'Maria Lima' '11888880001'
New-Patient 'Carla Rocha' 'carla.paciente@mindtrace.com' '12345678906' $false '1990-11-08T00:00:00Z' '11777770003' '' ''
New-Patient 'Diego Fernandes' 'diego.paciente@mindtrace.com' '12345678907' $false '1992-05-30T00:00:00Z' '11777770004' '' ''
New-Patient 'Elena Gomes' 'elena.paciente@mindtrace.com' '12345678908' $false '1997-09-12T00:00:00Z' '11777770005' '' ''
New-Patient 'Fabio Alves' 'fabio.paciente@mindtrace.com' '12345678909' $true '2012-01-25T00:00:00Z' '11777770006' 'Roberto Alves' '11888880002'
New-Patient 'Gabriela Pereira' 'gabriela.paciente@mindtrace.com' '12345678910' $false '1994-12-03T00:00:00Z' '11777770007' '' ''
New-Patient 'Henrique Souza' 'henrique.paciente@mindtrace.com' '12345678911' $true '2011-06-18T00:00:00Z' '11777770008' 'Patricia Souza' '11888880003'
New-Patient 'Isabela Martins' 'isabela.paciente@mindtrace.com' '12345678912' $false '1996-08-27T00:00:00Z' '11777770009' '' ''
New-Patient 'João Carvalho' 'joao.paciente@mindtrace.com' '12345678913' $true '2013-04-10T00:00:00Z' '11777770010' 'Fernanda Carvalho' '11888880004'
New-Patient 'Karla Ribeiro' 'karla.paciente@mindtrace.com' '12345678914' $false '1993-10-05T00:00:00Z' '11777770011' '' ''
New-Patient 'Lucas Teixeira' 'lucas.paciente@mindtrace.com' '12345678915' $false '1998-02-14T00:00:00Z' '11777770012' '' ''
New-Patient 'Mariana Santos' 'mariana.paciente@mindtrace.com' '12345678916' $false '1991-07-29T00:00:00Z' '11777770013' '' ''
New-Patient 'Nicolas Oliveira' 'nicolas.paciente@mindtrace.com' '12345678917' $true '2014-11-21T00:00:00Z' '11777770014' 'Sandra Oliveira' '11888880005'
New-Patient 'Olivia Silva' 'olivia.paciente@mindtrace.com' '12345678918' $false '1995-12-08T00:00:00Z' '11777770015' '' ''
New-Patient 'Paulo Mendes' 'paulo.paciente@mindtrace.com' '12345678922' $false '1989-03-17T00:00:00Z' '11777770016' '' ''
New-Patient 'Rafaela Nunes' 'rafaela.paciente@mindtrace.com' '12345678923' $false '1996-05-23T00:00:00Z' '11777770017' '' ''
New-Patient 'Sofia Castro' 'sofia.paciente@mindtrace.com' '12345678924' $true '2015-09-14T00:00:00Z' '11777770018' 'Amanda Castro' '11888880006'
New-Patient 'Thiago Barbosa' 'thiago.paciente@mindtrace.com' '12345678925' $false '1993-11-30T00:00:00Z' '11777770019' '' ''
New-Patient 'Vanessa Dias' 'vanessa.paciente@mindtrace.com' '12345678926' $false '1994-06-19T00:00:00Z' '11777770020' '' ''

Write-Host 'Seeding completed.' -ForegroundColor Green
