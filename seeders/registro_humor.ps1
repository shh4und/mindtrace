# PowerShell script to mimic registro_humor.sh behavior on Windows
$ErrorActionPreference = 'Stop'
Set-StrictMode -Version Latest

$ApiBase = 'http://localhost:8181/api/v1'
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
    if  (($response | Get-Member -MemberType All -Name "token")) { return $response.token }
    return $null
}

function New-HumorRecord {
    param (
        [string]$Token,
        [int]$NivelHumor,
        [int]$HorasSono,
        [int]$NivelStress,
        [int]$NivelEnergia,
        [string]$AutoCuidado,
        [string]$Observacoes,
        [datetime]$DataHora
    )

    $headers = @{ Authorization = "Bearer $Token" }
    $payload = @{
        nivel_humor      = $NivelHumor
        horas_sono       = $HorasSono
        nivel_stress     = $NivelStress
        nivel_energia    = $NivelEnergia
        auto_cuidado     = $AutoCuidado
        observacoes      = $Observacoes
        data_hora_registro = $DataHora.ToUniversalTime().ToString('yyyy-MM-ddTHH:mm:ssZ')
    } | ConvertTo-Json -Compress

    try {
        return Invoke-RestMethod -Method Post -Uri "$ApiBase/registro-humor/" -Headers $headers -ContentType 'application/json' -Body $payload -ErrorAction Stop
    } catch {
        Write-Warning "Record creation failed: $_"
        return $null
    }
}

function Get-Report {
    param (
        [string]$Token,
        [int]$Periodo
    )

    $headers = @{ Authorization = "Bearer $Token" }
    $uri = "$ApiBase/relatorios/?periodo=$Periodo"

    try {
        return Invoke-RestMethod -Method Get -Uri $uri -Headers $headers -ContentType 'application/json' -ErrorAction Stop
    } catch {
        Write-Warning "Report generation failed: $_"
        return $null
    }
}

Write-Host 'Starting humor registrations...' -ForegroundColor Cyan

for ($i = 0; $i -lt 4; $i++) {
    $pacEmail = $Pacientes[$i]
    $pacName = ($pacEmail.Split('@')[0]).Replace('.', ' ')

    Write-Host "Patient: $pacName ($pacEmail)" -ForegroundColor Yellow

    $pacToken = Invoke-ApiLogin -Email $pacEmail -Senha 'Password123!'
    if (-not $pacToken) {
        Write-Warning 'Login failed; skipping patient.'
        continue
    }

    $humorRecord = New-HumorRecord -Token $pacToken -NivelHumor 5 -HorasSono 8 -NivelStress 2 -NivelEnergia 8 -AutoCuidado 'Nenhum' -Observacoes "Registro de humor diario de $pacName" -DataHora (Get-Date)

    if ($humorRecord) {
        Write-Host 'Humor record created successfully.'
    } else {
        Write-Warning 'Failed to create humor record.'
        continue
    }

    $report = Get-Report -Token $pacToken -Periodo 7
    if ($report) {
        Write-Host 'Report data:'
        $report | ConvertTo-Json -Depth 6
    } else {
        Write-Warning 'Failed to retrieve report.'
    }

    Write-Host '---'
}

Write-Host 'Process completed.' -ForegroundColor Green
