# PowerShell script to mimic vincular.sh behavior on Windows
$ErrorActionPreference = 'Stop'
Set-StrictMode -Version Latest

$ApiBase = 'http://localhost:8181/api/v1'
$Profissionais = @(
    'joao.silva@mindtrace.com',
    'maria.santos@mindtrace.com',
    'pedro.oliveira@mindtrace.com',
    'ana.costa@mindtrace.com',
    'carlos.lima@mindtrace.com'
)
$Pacientes = @(
    'ana.paciente@mindtrace.com',
    'bruno.paciente@mindtrace.com',
    'carla.paciente@mindtrace.com',
    'diego.paciente@mindtrace.com',
    'elena.paciente@mindtrace.com',
    'fabio.paciente@mindtrace.com',
    'gabriela.paciente@mindtrace.com',
    'henrique.paciente@mindtrace.com',
    'isabela.paciente@mindtrace.com',
    'joao.paciente@mindtrace.com',
    'karla.paciente@mindtrace.com',
    'lucas.paciente@mindtrace.com',
    'mariana.paciente@mindtrace.com',
    'nicolas.paciente@mindtrace.com',
    'olivia.paciente@mindtrace.com',
    'paulo.paciente@mindtrace.com',
    'rafaela.paciente@mindtrace.com',
    'sofia.paciente@mindtrace.com',
    'thiago.paciente@mindtrace.com',
    'vanessa.paciente@mindtrace.com'
)

function Invoke-ApiLogin {
    param (
        [string]$Email,
        [string]$Senha
    )

    $payload = @{ email = $Email; senha = $Senha } | ConvertTo-Json -Compress

    try {
        $response = Invoke-RestMethod -Method Post -Uri "$ApiBase/entrar/login" -ContentType 'application/json' -Body $payload -ErrorAction Stop
    } catch {
        Write-Warning "Login failed for ${Email}: $_"
        return $null
    }

    if (($response | Get-Member -MemberType All -Name "dados") -and ($response | Get-Member -MemberType All -Name "dados.token")) { return $response.dados.token }
    if (($response | Get-Member -MemberType All -Name "token")) { return $response.token }
    return $null
}

function New-Invite {
    param (
        [string]$Token,
        [string]$PacienteEmail
    )

    $headers = @{ Authorization = "Bearer $Token" }
    $payload = @{ email_paciente = $PacienteEmail } | ConvertTo-Json -Compress

    try {
        $response = Invoke-RestMethod -Method Post -Uri "$ApiBase/convites/gerar" -Headers $headers -ContentType 'application/json' -Body $payload -ErrorAction Stop
        if (($response | Get-Member -MemberType All -Name "dados") -and ($response | Get-Member -MemberType All -Name "dados.token")) { return $response.dados.token }
        if  (($response | Get-Member -MemberType All -Name "token")) { return $response.token }
    } catch {
        Write-Warning "Invite generation failed for ${PacienteEmail}: $_"
    }

    return $null
}

function Invoke-LinkPatient {
    param (
        [string]$TokenPaciente,
        [string]$TokenConvite
    )

    $headers = @{ Authorization = "Bearer $TokenPaciente" }
    $payload = @{ token = $TokenConvite } | ConvertTo-Json -Compress

    try {
        return Invoke-RestMethod -Method Post -Uri "$ApiBase/convites/vincular" -Headers $headers -ContentType 'application/json' -Body $payload -ErrorAction Stop
    } catch {
        Write-Warning "Linking failed: $_"
        return $null
    }
}

Write-Host 'Starting patient linking process...' -ForegroundColor Cyan

$pacienteIndex = 0

for ($i = 0; $i -lt $Profissionais.Count; $i++) {
    $profEmail = $Profissionais[$i]
    $profNome = ($profEmail.Split('@')[0]).Replace('.', ' ')

    Write-Host "Professional: $profNome ($profEmail)" -ForegroundColor Yellow

    $profToken = Invoke-ApiLogin -Email $profEmail -Senha 'Password123!'
    if (-not $profToken) {
        Write-Warning 'Professional login failed; skipping.'
        continue
    }

    for ($j = 1; $j -le 4 -and $pacienteIndex -lt $Pacientes.Count; $j++) {
        $pacEmail = $Pacientes[$pacienteIndex]
        $pacNome = ($pacEmail.Split('@')[0]).Replace('.', ' ')
        Write-Host "  Patient $j/4: $pacNome"

        $conviteToken = New-Invite -Token $profToken -PacienteEmail $pacEmail
        if (-not $conviteToken) {
            Write-Warning '  Failed to create invite.'
            $pacienteIndex++
            continue
        }

        $pacToken = Invoke-ApiLogin -Email $pacEmail -Senha 'Password123!'
        if (-not $pacToken) {
            Write-Warning '  Patient login failed.'
            $pacienteIndex++
            continue
        }

        $resultado = Invoke-LinkPatient -TokenPaciente $pacToken -TokenConvite $conviteToken
        if ($resultado -and ($resultado.mensagem -match 'sucesso' -or $resultado.mensagem -match 'vinculad')) {
            Write-Host '  Patient linked successfully.'
        } elseif ($resultado) {
            Write-Host "  API response: $($resultado | ConvertTo-Json -Depth 4)"
        } else {
            Write-Warning '  Linking call failed.'
        }

        $pacienteIndex++
    }
}

Write-Host 'Linking process completed.' -ForegroundColor Green
