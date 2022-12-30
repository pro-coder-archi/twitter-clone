package com.twitterclone.theme

import androidx.compose.ui.graphics.Color

private val primaryColor= Color(0xFF1D9BF0)
private val backgroundColor= Color(0XFF000000)
private val contrastColor= Color(0xFFFFFFFF)
private val errorColor= Color(0xFFFF453A)
private val textColor= Color.White.copy(alpha= 0.5f)
private val lightBorderColor= Color.White.copy(alpha = 0.2f)

class ThemeColors(

    var primary: Color,
    var background: Color,
    var error: Color,
    var text: Color,
    var lightBorder: Color,
    var contrast: Color,

    var isForDarkTheme: Boolean
)

fun getDarkThemeColors(

    primary: Color= primaryColor,
    background: Color= backgroundColor,
    contrast: Color= contrastColor,
    error: Color= errorColor,
    text: Color= textColor,
    lightBorder: Color= lightBorderColor

) =
    ThemeColors(
        primary= primary,
        background= background,
        contrast= contrast,
        error= error,
        text = text,
        lightBorder= lightBorder,

        isForDarkTheme= true
    )