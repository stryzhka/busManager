import {Button, MenuListItem, ScrollView, styleReset, TextInput} from "react95";
import React from "react";
import {createGlobalStyle} from "styled-components";
import ms_sans_serif from "../assets/fonts/fixedsys.woff2";
import ms_sans_serif_bold from 'react95/dist/fonts/ms_sans_serif_bold.woff2';
import original from 'react95/dist/themes/original';

const GlobalStyles = createGlobalStyle`
  ${styleReset}
  @font-face {
    font-family: 'ms_sans_serif';
    src: url('${ms_sans_serif}') format('woff2');
    font-weight: 400;
    font-style: normal
  }
  @font-face {
    font-family: 'ms_sans_serif';
    src: url('${ms_sans_serif_bold}') format('woff2');
    font-weight: bold;
    font-style: normal
  }
  body, input, select, textarea {
    font-family: 'ms_sans_serif';
  }
`;

const DriverComponent = () => {
    return (
        <div style={{ display: 'flex', flexDirection: 'row', alignItems: 'flex-start' }}>
            <ScrollView style={{width: '200px', height: '200px', marginRight: '20px'}}>
                <MenuListItem>хуй</MenuListItem>
                <MenuListItem>долбоеб</MenuListItem>
                <MenuListItem>Быков андрей еве</MenuListItem>
                <MenuListItem>penus</MenuListItem>
                <MenuListItem>ВАДЫЛА</MenuListItem>
                <MenuListItem>loroo e0id asui</MenuListItem>

            </ScrollView>
            <div>
                <div style={{ display: 'flex', flexDirection: 'row', alignItems: 'flex-start', marginBottom: '10px' }}>
                    <TextInput style={{width: '150px', marginRight: '20px'}}></TextInput>
                    <div style={{marginRight: '20px'}}>Имя</div>
                </div>
                <div style={{ display: 'flex', flexDirection: 'row', alignItems: 'flex-start', marginBottom: '10px' }}>
                    <TextInput style={{width: '150px', marginRight: '20px'}}></TextInput>
                    <div style={{marginRight: '20px'}}>Фамилия</div>
                </div>
                <div style={{ display: 'flex', flexDirection: 'row', alignItems: 'flex-start', marginBottom: '10px' }}>
                    <TextInput style={{width: '150px', marginRight: '20px'}}></TextInput>
                    <div style={{marginRight: '20px'}}>Отчество</div>
                </div>
                <div style={{ display: 'flex', flexDirection: 'row', alignItems: 'flex-start', marginBottom: '10px' }}>
                    <TextInput style={{width: '150px', marginRight: '20px'}}></TextInput>
                    <div style={{marginRight: '20px'}}>Дата рождения</div>
                </div>
                <div style={{ display: 'flex', flexDirection: 'row', alignItems: 'flex-start', marginBottom: '10px' }}>
                    <TextInput style={{width: '170px', marginRight: '20px'}}></TextInput>
                    <div style={{marginRight: '20px'}}>Серия и номер паспорта</div>
                </div>
                <div style={{ display: 'flex', flexDirection: 'row', alignItems: 'flex-start', marginBottom: '10px' }}>
                    <TextInput style={{width: '150px', marginRight: '20px'}}></TextInput>
                    <div style={{marginRight: '20px'}}>Номер СНИЛС</div>
                </div>
                <div style={{ display: 'flex', flexDirection: 'row', alignItems: 'flex-start',  }}>
                    <Button style={{marginRight: '10px'}}>Сохранить</Button>
                    <Button style={{marginRight: '10px'}}>Удалить</Button>
                    <Button >Создать</Button>
                </div>
            </div>

        </div>
    )
}

export default DriverComponent