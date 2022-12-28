package com.example.twitterclone.theme

import androidx.compose.runtime.Composable
import androidx.compose.runtime.CompositionLocalProvider
import androidx.compose.runtime.ReadOnlyComposable

object AppTheme {

    val colors: ThemeColors
        get( )= getDarkThemeColors( )

    val typography: ThemeTypography
        get( )= ThemeTypography( )

}

// TODO: implement light theme and functionality to switch between themes